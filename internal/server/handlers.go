package server
import("encoding/json";"hash/fnv";"net/http";"strconv";"github.com/stockyard-dev/stockyard-drifter/internal/store")
func(s *Server)handleListFlags(w http.ResponseWriter,r *http.Request){list,_:=s.db.ListFlags();if list==nil{list=[]store.Flag{}};writeJSON(w,200,list)}
func(s *Server)handleGetFlag(w http.ResponseWriter,r *http.Request){key:=r.PathValue("key");f,_:=s.db.GetFlag(key);if f==nil{writeError(w,404,"not found");return};writeJSON(w,200,f)}
func(s *Server)handleCreateFlag(w http.ResponseWriter,r *http.Request){
    if !s.limits.IsPro(){n,_:=s.db.CountFlags();if n>=10{writeError(w,403,"free tier: 10 flags max");return}}
    var f store.Flag;json.NewDecoder(r.Body).Decode(&f)
    if f.Key==""{writeError(w,400,"key required");return}
    if f.Rollout==0{f.Rollout=100}
    if err:=s.db.CreateFlag(&f);err!=nil{writeError(w,500,err.Error());return}
    writeJSON(w,201,f)}
func(s *Server)handleUpdateFlag(w http.ResponseWriter,r *http.Request){
    key:=r.PathValue("key");existing,_:=s.db.GetFlag(key);if existing==nil{writeError(w,404,"not found");return}
    json.NewDecoder(r.Body).Decode(existing);existing.Key=key
    if err:=s.db.UpdateFlag(existing);err!=nil{writeError(w,500,err.Error());return}
    writeJSON(w,200,existing)}
func(s *Server)handleDeleteFlagByKey(w http.ResponseWriter,r *http.Request){key:=r.PathValue("key");f,_:=s.db.GetFlag(key);if f!=nil{s.db.DeleteFlag(f.ID)};writeJSON(w,200,map[string]string{"status":"deleted"})}
func(s *Server)handleToggleFlag(w http.ResponseWriter,r *http.Request){
    key:=r.PathValue("key");f,_:=s.db.GetFlag(key);if f==nil{writeError(w,404,"not found");return}
    f.Enabled=!f.Enabled;s.db.UpdateFlag(f);writeJSON(w,200,map[string]interface{}{"key":key,"enabled":f.Enabled})}
func evalFlag(f *store.Flag,userID string)bool{
    if !f.Enabled{return false}
    if f.Rollout>=100{return true}
    if userID==""{return false}
    h:=fnv.New32a();h.Write([]byte(f.Key+":"+userID))
    return int(h.Sum32()%100)<f.Rollout}
func(s *Server)handleEval(w http.ResponseWriter,r *http.Request){
    key:=r.URL.Query().Get("key");userID:=r.URL.Query().Get("user_id")
    f,_:=s.db.GetFlag(key);if f==nil{writeJSON(w,200,map[string]interface{}{"key":key,"enabled":false,"reason":"not_found"});return}
    result:=evalFlag(f,userID);go s.db.LogEval(key,userID,result)
    writeJSON(w,200,map[string]interface{}{"key":key,"enabled":result,"rollout":f.Rollout})}
func(s *Server)handleEvalAll(w http.ResponseWriter,r *http.Request){
    userID:=r.URL.Query().Get("user_id")
    flags,_:=s.db.ListFlags();out:=map[string]bool{}
    for _,f:=range flags{out[f.Key]=evalFlag(&f,userID)}
    writeJSON(w,200,out)}
func(s *Server)handleStats(w http.ResponseWriter,r *http.Request){f,_:=s.db.CountFlags();e,_:=s.db.CountEvals();writeJSON(w,200,map[string]interface{}{"flags":f,"evals":e})}
var _ = strconv.Itoa

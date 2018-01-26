package main;
import (
  "io"
  "os"
  "fmt"
  "strconv"
  "regexp"
  "net/http"
  "muen"
  "./view"
  "./woo"
);

func StrToInt(s string, def int) int{
  ret, err:=strconv.Atoi(s);
  if(err!=nil){
    return def;
  }
  return ret;
}

func main() {
  const Header_Server = "ESSAY/1.1";
  //ch := make(chan int);
  woo.ResetMain();
  os.MkdirAll("./www/home/", os.ModePerm);
  os.MkdirAll("./www/common/", os.ModePerm);  
  os.MkdirAll("./www/var/", os.ModePerm);  

  http.Handle("/home/", http.FileServer(http.Dir("./www/")));
  http.Handle("/common/", http.FileServer(http.Dir("./www/")));
  http.Handle("/var/", http.FileServer(http.Dir("./www/")));
  http.HandleFunc("/now.do", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Server", Header_Server);
    w.Write([]byte(muen.Now()));
  });
  http.HandleFunc("/about.do", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Server", Header_Server);
    io.WriteString(w, woo.AboutMe());
  });
  http.HandleFunc("/save.do", func(w http.ResponseWriter, r *http.Request) {
    method:= r.Method;
    if(method=="POST"){
      rowid:= woo.Save(r.FormValue("content"));
      w.Header().Set("Server", Header_Server);
      w.Header().Set("Location", "/view/" + rowid + ".html");
      w.WriteHeader(302); 
      io.WriteString(w, `{ret: true, rowid:"` + rowid + `"}`);
      muen.Sendln(`[ESSAY] NEW: rowid=` + rowid);
    } else {
      io.WriteString(w, "Method: POST;\r\n");
      io.WriteString(w, "Parametes: content;\r\n");
    }
  });
  http.HandleFunc("/load.do", func(w http.ResponseWriter, r *http.Request) {
    rowid:= r.FormValue("rowid");
    w.Header().Set("Server", Header_Server);
    w.Header().Set("Content-Type", "text/plain; charset=utf-8");
    io.WriteString(w, woo.Load(rowid));
  });
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Server", Header_Server);
    w.Header().Set("Content-Type", "text/html; charset=utf-8");  
    io.WriteString(w, view.HtmlIndex());
  });
  http.HandleFunc("/view/", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Server", Header_Server);
    w.Header().Set("Content-Type", "text/html; charset=utf-8");
    reg := regexp.MustCompile(`^/view/(\d{4}\d{2}\d{2}\d{40}).html`);
    mc:= reg.FindStringSubmatch(r.URL.Path);   
    if(len(mc)>0){
      rowid:= mc[1];
      io.WriteString(w, view.HtmlView(rowid, woo.Load(rowid)));    
    } else {
      io.WriteString(w, view.HtmlError("BAD-ROWID"));      
    }
  });  
  fmt.Println(" ");
  fmt.Println(" ESSAY-SERVER: http://127.0.0.1:13800");
  fmt.Println(" EMAIL: woo@omuen.com");
  fmt.Println(" SERVER-UUID: " + woo.GetServerId());
  fmt.Println("------------------------------------------------------------------");
  fmt.Println(" Start at " + muen.Now());
  muen.Send("<!>\r\n");
  muen.Sendln("[ESSAY] Start at " + muen.Now());
  http.ListenAndServe(":13800", nil);
}
package woo;
import (
  "os"
  "fmt"
  "regexp"
  "database/sql"
  "muen"
  _ "github.com/mattn/go-sqlite3"
);

func ResetMain(){
  dbname:= "./main.db";
  _, err := os.Stat(dbname);
  if(!(err == nil || os.IsExist(err))){
    db, _ := sql.Open("sqlite3", dbname);
    uuid:= muen.NewKey();
    cmd:=`
      CREATE TABLE IF NOT EXISTS [main] ([rowid] PRIMARY KEY, [content]);
      CREATE TABLE IF NOT EXISTS [meta] ([rowid] PRIMARY KEY, [name], [content]);
      INSERT INTO [meta] ([rowid], [name], [content]) VALUES('about.author', 'SYSTEM', '杨波(BambooYoung)');
      INSERT INTO [meta] ([rowid], [name], [content]) VALUES('about.email', 'SYSTEM', 'lokme@foxmail.com');
      INSERT INTO [meta] ([rowid], [name], [content]) VALUES('about.website', 'SYSTEM', 'https://www.bimwook.com');
      INSERT INTO [meta] ([rowid], [name], [content]) VALUES('server.uuid', 'SYSTEM', ?);
      INSERT INTO [meta] ([rowid], [name], [content]) VALUES('server.created', 'SYSTEM', ?);
      INSERT INTO [main] ([rowid], [content]) VALUES('server.uuid', ?);
    `;
    _, err := db.Exec(cmd, uuid, muen.Now(), uuid);
    if err!=nil {
      fmt.Println(err);
    }
    defer db.Close();
  }
}

func GetServerId() string{
  ResetMain();
  db, _ := sql.Open("sqlite3", "./main.db");
  uuid:= "BAD-KEY";
  cmd:=`SELECT [content] FROM [main] WHERE [rowid]='server.uuid';`;
  rows, err := db.Query(cmd);
  if (err!=nil) {
    fmt.Println(err);
  } else {
    for rows.Next() {
      var content string;
      if err:= rows.Scan(&content); err == nil {
        uuid = content;
      }
    }
  }
  defer db.Close(); 
  return uuid;
}

func GetCache(rowid string) (string, bool) {
  reg := regexp.MustCompile(`^(\d{4})(\d{2})(\d{2})\d{40}$`);
  fas:= reg.FindStringSubmatch(rowid);
  ret:= "";
  if(len(fas)>0){
    ret = "./cache/at" + fas[1] + "/m" + fas[2] + "/d" + fas[3] + ".db";
  }
  _, err := os.Stat(ret);
  return ret, err == nil || os.IsExist(err) ;
}

func ResetCache(rowid string) bool{
  reg := regexp.MustCompile(`^(\d{4})(\d{2})(\d{2})\d{40}$`);
  fas:= reg.FindStringSubmatch(rowid);
  if(len(fas)==0){
    return false;
  }
  fn:= "./cache/at" + fas[1] + "/m" + fas[2];
  os.MkdirAll(fn, os.ModePerm);
  dbname, _:= GetCache(rowid);
  db, _ := sql.Open("sqlite3", dbname);
  cmd:=`
    CREATE TABLE IF NOT EXISTS [meta] ([rowid] PRIMARY KEY, [name], [content]);
    CREATE TABLE IF NOT EXISTS [main] ([rowid] PRIMARY KEY, [uid], [ua], [content], [ext], [read], [rank], [created]);
    INSERT INTO [meta] ([rowid], [name], [content]) VALUES('server.uuid', 'SYSTEM', ?);
    INSERT INTO [meta] ([rowid], [name], [content]) VALUES('db.uuid', 'SYSTEM', ?);
    INSERT INTO [meta] ([rowid], [name], [content]) VALUES('db.created', 'SYSTEM', ?);
  `;
  _, err := db.Exec(cmd, GetServerId(), muen.NewKey(), muen.Now());
  if err!=nil {
    fmt.Println(err);
  }
  defer db.Close();  
  return true;
}

func Save(data string) string {
  rowid := muen.NewKey();
  fn,exists := GetCache(rowid);
  if(!exists){
    ResetCache(rowid);
  }
  db, _ := sql.Open("sqlite3", fn);
  cmd:=`
    INSERT INTO [main] ([rowid], [uid], [ua], [content], [ext], [read], [rank], [created]) VALUES(?,?,?,?,?,?,?,?);
  `;
  _, err := db.Exec(cmd, rowid, "", "", data, "", 0,0, muen.Now());
  if err!=nil {
    fmt.Println(err);
  }
  defer db.Close();
  return rowid;
}

func Load(rowid string) string {
  fn,exists := GetCache(rowid);
  if(exists){
    var data string;
    db, _ := sql.Open("sqlite3", fn);
    cmd:=`SELECT [content] FROM [main] WHERE [rowid]=?;`;
    rows, err := db.Query(cmd, rowid);
    if (err!=nil) {
      fmt.Println(err);
    } else {
      for rows.Next() {
        var content string;
        if err:= rows.Scan(&content); err == nil {
          data = content;
        }else{
          data = "null"
        }
      }
    }
    defer rows.Close();
    defer db.Close();  
    return data;
  }
  return "BAD-ROWID";
}
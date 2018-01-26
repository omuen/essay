package woo;
import (
  "bytes"
);

func AboutMe() string{
  var buf bytes.Buffer;
  buf.WriteString(`Name: 杨波 (Bamboo Young)` + "\r\n");
  buf.WriteString(`QQ: 3262706` + "\r\n");
  buf.WriteString(`E-Mail: woo@omuen.com` + "\r\n");
  buf.WriteString(`Memo: OK` + "\r\n");
  return buf.String();  
}
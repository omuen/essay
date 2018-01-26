package view;
import (
  "bytes"
  "muen"
);

func HtmlError(data string) string{
  var buf bytes.Buffer;
  buf.WriteString(`<!DOCTYPE html>` + "\r\n");
  buf.WriteString(`<html>` + "\r\n");
  buf.WriteString(`  <meta charset="utf-8">` + "\r\n");
  buf.WriteString(`  <title>Hello</title>` + "\r\n");
  buf.WriteString(`  <body>` + "\r\n");
  buf.WriteString(`    ` + muen.HtmlEncode(data) + "\r\n");
  buf.WriteString(`  </body>` + "\r\n");
  buf.WriteString(`</html>` + "\r\n");
  return buf.String();
}

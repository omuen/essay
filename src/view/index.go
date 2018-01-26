package view;
import (
  "bytes"
);

func HtmlIndex() string{
  var buf bytes.Buffer;
  buf.WriteString(`  <!DOCTYPE html>` + "\r\n");
  buf.WriteString(`<html lang="zh-CN">` + "\r\n");
  buf.WriteString(`  <head>` + "\r\n");
  buf.WriteString(`    <meta charset="utf-8" />` + "\r\n");
  buf.WriteString(`    <meta content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no" name="viewport" />` + "\r\n");
  buf.WriteString(`    <title></title>` + "\r\n");
  buf.WriteString(`    <style type="text/css">` + "\r\n");
  buf.WriteString(`      html{margin:0; padding:0;}` + "\r\n");
  buf.WriteString(`      body{margin:0; padding:4px;}` + "\r\n");
  buf.WriteString(`      textarea{` + "\r\n");
  buf.WriteString(`        width: 100%;` + "\r\n");
  buf.WriteString(`        max-width: 640px;` + "\r\n");
  buf.WriteString(`        display: block;` + "\r\n");
  buf.WriteString(`        resize: none;` + "\r\n");
  buf.WriteString(`        outline: none;` + "\r\n");
  buf.WriteString(`        height: 240px;` + "\r\n");
  buf.WriteString(`        margin: 0 auto;` + "\r\n");
  buf.WriteString(`        box-sizing: border-box;` + "\r\n");
  buf.WriteString(`        border: solid 1px #EAEAEA;` + "\r\n");
  buf.WriteString(`        padding: 8px;` + "\r\n");
  buf.WriteString(`        line-height: 1.618;` + "\r\n");
  buf.WriteString(`        text-align: left;` + "\r\n");
  buf.WriteString(`      }` + "\r\n");
  buf.WriteString(`      input[type="submit"]{` + "\r\n");
  buf.WriteString(`        color: white;` + "\r\n");
  buf.WriteString(`        width: 100%;` + "\r\n");
  buf.WriteString(`        max-width: 640px;` + "\r\n");
  buf.WriteString(`        display: block;` + "\r\n");
  buf.WriteString(`        resize: none;` + "\r\n");
  buf.WriteString(`        outline: none;` + "\r\n");
  buf.WriteString(`        border: none;` + "\r\n");
  buf.WriteString(`        border-radius: 5px;` + "\r\n");
  buf.WriteString(`        padding: 8px;` + "\r\n");
  buf.WriteString(`        margin: 10px auto;` + "\r\n");
  buf.WriteString(`        box-sizing: border-box;` + "\r\n");
  buf.WriteString(`        background-color: #008080;` + "\r\n");
  buf.WriteString(`      }` + "\r\n");
  buf.WriteString(`    </style>` + "\r\n");
  buf.WriteString(`  </head>` + "\r\n");
  buf.WriteString(`  <body>` + "\r\n");
  buf.WriteString(`    <form  action="/save.do" target="_blank" method="POST">` + "\r\n");
  buf.WriteString(`      <textarea name="content"></textarea>` + "\r\n");
  buf.WriteString(`      <input type="submit" value="Save" />` + "\r\n");
  buf.WriteString(`    </form>` + "\r\n");
  buf.WriteString(`  </body>` + "\r\n");
  buf.WriteString(`</html>` + "\r\n");
  return buf.String();
}

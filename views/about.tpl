<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>{{.pro.Realname}}-{{config "String" "globaltitle" ""}}</title>
{{template "inc/meta.tpl" .}}
</head>
<body>
{{template "inc/head.tpl" .}}
<div class="container">
  <div class="row">
    <div class="container">
      <div class="page-header">
        <h1>{{.pro.Realname}}</h1>
      </div>
      <div class="jumbotron">
        <p>{{if .pro.Sex}}超级Man{{else}}小小女生{{end}}</p>
        <p>出生：{{.pro.Birth}}</p>
        <p>邮箱是：{{.pro.Email}}</p>
        <p>这是我的电话： {{.pro.Phone}}</p>
        <p>我的地址： {{.pro.Address}}</p>
        <p>经常会做这些事：{{.pro.Hobby}}</p>
        <p>我的介绍：{{.pro.Intro}}</p>
      </div>
    </div>
  </div>
</div>
</body>
</html>

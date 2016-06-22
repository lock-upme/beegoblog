<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>{{.art.Title}}-{{config "String" "globaltitle" ""}}</title>
<meta name="keywords" content="{{.art.Title}}">
<meta name="description" content="{{.art.Summary}}">
{{template "inc/meta.tpl" .}}
</head>
<body>
{{template "inc/head.tpl" .}}
<div class="container">
  <div class="row">
    <div class="container">
      <div class="page-header">
        <h1>{{.art.Title}}
          <div class="bdsharebuttonbox pull-right"><a href="#" class="bds_more" data-cmd="more"></a><a href="#" class="bds_qzone" data-cmd="qzone" title="分享到QQ空间"></a><a href="#" class="bds_tsina" data-cmd="tsina" title="分享到新浪微博"></a><a href="#" class="bds_weixin" data-cmd="weixin" title="分享到微信"></a></div>
        </h1>
      </div>
      <div class="lead">{{.art.Summary}}</div>
      <div>{{str2html .art.Content}}</div>
    </div>
    <div class="container"> {{if .coms}}
      <h3 class="alert-success padtb10">评论
        <div class="pull-right"><a href="#comment">我要评论</a></div>
      </h3>
      {{range $k,$v := .coms}}
      <div class="media">
        <div class="media-left"> <a href="{{$v.Uri}}" title="{{$v.Nickname}}" target="_blank"> <img class="media-object" alt="{{$v.Nickname}}" style="width: 64px; height: 64px;" src="{{avatar}}" data-holder-rendered="true"> </a> </div>
        <div class="media-body">
          <h4 class="media-heading">{{$v.Nickname}} <span class="label label-info pull-right">（{{date_mh $v.Created}}）</span></h4>
          <div> {{$v.Content}}
            {{if $.isLogin }}
            <p class="pull-right"> {{if eq $v.Status 0}} <a href="javascript:;" class="btn btn-danger btn-xs js-comment-status" data-status="1" data-id="{{$v.Id}}">正常</a> {{else}} <a href="javascript:;" class="btn btn-danger btn-xs js-comment-status" data-status="0" data-id="{{$v.Id}}">屏蔽</a> {{end}} </p>
            {{end}} </div>
        </div>
      </div>
      {{end}}
      {{end}}
      
      {{if .paginator.HasPages}}
      <ul class="pagination pagination">
        {{if .paginator.HasPrev}}
        <li><a href="{{.paginator.PageLinkFirst}}">首页</a></li>
        <li><a href="{{.paginator.PageLinkPrev}}">&laquo;</a></li>
        {{else}}
        <li class="disabled"><a>首页</a></li>
        <li class="disabled"><a>&laquo;</a></li>
        {{end}}
        {{range $index, $page := .paginator.Pages}} <li{{if $.paginator.IsActive .}} class="active"{{end}}> <a href="{{$.paginator.PageLink $page}}">{{$page}}</a>
        </li>
        {{end}}
        {{if .paginator.HasNext}}
        <li><a href="{{.paginator.PageLinkNext}}">&raquo;</a></li>
        <li><a href="{{.paginator.PageLinkLast}}">尾页</a></li>
        {{else}}
        <li class="disabled"><a>&raquo;</a></li>
        <li class="disabled"><a>尾页</a></li>
        {{end}}
      </ul>
      {{end}}
      <h3 class="alert-success padtb10">我要评论<a name="comment"></a></h3>
      <form method="post" id="comment-form">
        <div class="form-group">
          <label for="nickname">昵称</label>
          <input type="text" class="form-control" id="nickname" name="nickname" placeholder="nickname" value="">
        </div>
        <div class="form-group">
          <label for="uri">URL</label>
          <input type="url" class="form-control" id="uri" name="uri" placeholder="http://" value="">
        </div>
        <div class="form-group">
          <label for="content">内容</label>
          <textarea class="form-control" id="content" name="content" placeholder="Content" style="height:120px;"></textarea>
        </div>
        <input type="hidden" name="article_id" value="{{.art.Id}}">
        <input type="hidden" name="id" value="{{.com.Id}}">
        <button type="submit" class="btn btn-primary">提交评论</button>
      </form>
    </div>
  </div>
</div>
{{template "inc/foot.tpl" .}}
<script>window._bd_share_config={"common":{"bdSnsKey":{},"bdText":"","bdMini":"2","bdMiniList":false,"bdPic":"","bdStyle":"0","bdSize":"32"},"share":{}};with(document)0[(getElementsByTagName('head')[0]||body).appendChild(createElement('script')).src='http://bdimg.share.baidu.com/static/api/js/share.js?v=89860593.js?cdnversion='+~(-new Date()/36e5)];</script>
</body>
</html>

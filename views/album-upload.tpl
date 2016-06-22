<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>图片上传-{{config "String" "globaltitle" ""}}</title>
{{template "inc/meta.tpl" .}}
<link href="/static/js/bootstrap-fileinput-master/css/fileinput.min.css" media="all" rel="stylesheet" type="text/css" />
</head>
<body>
{{template "inc/head.tpl" .}}
<div class="container">
  <div class="row">
    <form method="post" action="/uploadmulti" enctype="multipart/form-data"  id="uploadMulti-form">
       <div class="wrapper text-center">

            <h2>请选择图片</h2>
			<input id="albumUpload" name="uploadFiles" type="file" multiple class="file-loading" data-allowed-file-extensions='["jpg", "jpeg", "png", "gif"]'>

        </div>
    </form>
  </div>
</div>
{{template "inc/foot.tpl" .}}
<script src="/static/js/bootstrap-fileinput-master/js/plugins/canvas-to-blob.min.js" type="text/javascript"></script>
<script src="/static/js/bootstrap-fileinput-master/js/plugins/sortable.min.js" type="text/javascript"></script>
<script src="/static/js/bootstrap-fileinput-master/js/plugins/purify.min.js" type="text/javascript"></script>
<script src="/static/js/bootstrap-fileinput-master/js/fileinput.min.js"></script>
<script src="/static/js/bootstrap-fileinput-master/themes/fa/fa.js"></script>
<script src="/static/js/bootstrap-fileinput-master/js/locales/zh.js"></script>

<script>
$(function(){	
	$("#albumUpload").fileinput({ language: 'zh', showCaption: false});	
});
</script>
</body>
</html>

$(function(){
    
	//评论状态
	$('.js-comment-status').on('click',function(){
		var that = $(this);
		var status = that.attr('data-status');
		var id = that.attr('data-id');
	    $.post("/comment/edit/status", { id: id,status:status}, function(data){
			dialogInfo(data.message)
			if (data.code == 1) {				
				that.attr({'data-status':status == 1 ? 0 : 1}).text(status == 1? '屏蔽':'正常');
			} else {
				
			}
			setTimeout(function(){ $('#dialogInfo').modal('hide'); }, 1000);
		},'json');
	});
	
	
    //添加评论
	$('#comment-form').validate({
		ignore:'',
        rules : {
            nickname:{ required : true,rangelength : [1, 10]},
            content:{required : true}
        },
        messages : {
            nickname : {required : '请填写昵称',rangelength : '昵称最多10个字'},
            content : {required: '请填写评论内容'}
        },
        submitHandler:function(form) {
			var id = $('input[name="id"]').val();
			var url = id ? '/comment/edit/'+id : '/comment/add';
            $(form).ajaxSubmit({
                url:url,
                type:'POST',
                dataType:'json',
                success:function(data) {
                    dialogInfo(data.message)
                    if (data.code == 1) {
                      $('#comment-form')[0].reset();
                    } else {
                        
                    }
					setTimeout(function(){ $('#dialogInfo').modal('hide'); }, 1000);
                }
            });
        }
		
	});
    
	//添加及修改文章
    $('#article-form').validate({        
        ignore:'',
        rules : {
            title:{ required : true,rangelength : [1, 255]},
            content:{required : true}
        },
        messages : {
            title : {required : '博客标题必填',rangelength : '标题最多255字'},
            content : {required: '正文必填'}
        },
        submitHandler:function(form) {
			var id = $('input[name="id"]').val();
			var url = id != 0 ? '/article/edit/'+id : '/article/add';
            $(form).ajaxSubmit({
                url:url,
                type:'POST',
                dataType:'json',
                success:function(data) {
                    dialogInfo(data.message)
                    if (data.code == 1) {
						url = id ? '/article':'/article/add';
						setTimeout(function(){window.location.href=url}, 1000);
                    } else {
                        
                    }
                }
            });
        }
    });
	
	//多图上传	
	$('#uploadMulti-form').validate({        
        ignore:'',
        rules : {
            uploadFiles:{ required : true}
        },
        messages : {
            uploadFiles:{required : '请选择上传图片，可以多选'}
        },
        submitHandler:function(form) {
			var id = $('input[name="id"]').val();
			var url = '/uploadmulti';
            $(form).ajaxSubmit({
                url:url,
                type:'POST',
                dataType:'json',
                success:function(data) {
                    dialogInfo(data.message)
                    if (data.code == 1) {						
						setTimeout(function(){window.location.href='/album/'}, 1000);
                    } else {
                        
                    }
                }
            });
        }
    });

	//登录
    $('#login-form').validate({        
        ignore:'',
        rules : {
            phone:{ required : true, number:true,rangelength:[11,11]},
            password:{required : true}
        },
        messages : {
            phone : {required : '请填写手机号',number:'请填写正确的手机号',rangelength : '请填写11位手机号'},
            password : {required: '请填写密码'}
        },
        submitHandler:function(form) {
            var url = '/login';
            $(form).ajaxSubmit({
                url:url,
                type:'POST',
                dataType:'json',
                success:function(data) {
                    dialogInfo(data.message)
                    if (data.code == 1) {
                       setTimeout(function(){window.location.reload();}, 2000);
                    } else {
                        
                    }
                }
            });
        }
    });	
	

	$('body').delegate('.js-album-submit', 'click', function(){
		var that = $(this);
		var form = $('#album-form');
		var title = form.find('input[name="title"]');
		if (title.val() == '') { title.focus(); return; }
		var summary = form.find('textarea[name="summary"]');
		if (summary.val() == '') { summary.focus(); return; }
		var id = form.find('input[name="id"]');
		if (id.val() == '') { return; }
		
		$.post("/album/edit", { id: id.val(), title:title.val(), summary:summary.val()}, function(data){			
			if (data.code == 1) {				
				that.text(data.message)
			} else {
				dialogInfo(data.message)
			}
			setTimeout(function(){ $('#dialogInfo').modal('hide'); }, 1000);
		},'json');	
		
	});	
	
	$('.js-album-edit').on('click', function(){
		var that = $(this);
		var title = that.attr('data-title');
		var summary = that.attr('data-summary');
		var id = that.attr('data-id');
		var status = that.attr('data-status');
		dialogAlbum(id, title, summary, status);
	});
	
});

function dialogInfo(msg) {
	var html = '';
	html += '<div class="modal fade" id="dialogInfo" tabindex="-1" role="dialog" aria-labelledby="dialogInfoTitle">';
  	html += '<div class="modal-dialog" role="document">';
    html += '<div class="modal-content">';
    html += '<div class="modal-header">';
    html += '<button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>';
    html += '<h4 class="modal-title" id="dialogInfoTitle">信息提示</h4>';
    html += ' </div>';
    html += '<div class="modal-body">';
    html += '<p>'+msg+'</p>';
    html += '</div>';
    //html += '<div class="modal-footer">';
    //html += ' <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>';
    //html += ' <button type="button" class="btn btn-primary">Send message</button>';
    //html += '</div>';
    html += '</div>';
  	html += '</div>';
	html += '</div>';
	$('body').append(html);
	$('#dialogInfo').modal('show'); 
}

function dialogAlbum(id, title, summary, status) {
	$('#dialogAlbum').remove();
	var html = '';
	html += '<div class="modal fade in" id="dialogAlbum" tabindex="-1" role="dialog" aria-labelledby="dialogAlbumTitle">';
    html += '  <div class="modal-dialog" role="document">';
	html += '   <form id="album-form">';
    html += '    <div class="modal-content">';
    html += '      <div class="modal-header">';
    html += '        <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">×</span></button>';
    html += '        <h4 class="modal-title" id="dialogAlbumTitle">编辑'+title+'</h4>';
    html += '      </div>';
    html += '      <div class="modal-body">';
    html += '          <div class="form-group">';
    html += '            <label for="recipient-name" class="control-label">标题:</label>';
    html += '           <input type="text" class="form-control" name="title" value="'+title+'">';
    html += '          </div>';
	
    html += '          <div class="form-group">';
    html += '            <label for="message-text" class="control-label">说明:</label>';
    html += '            <textarea class="form-control" name="summary">'+summary+'</textarea>';
    html += '          </div>';
	
	
	html += '<div class="form-group">';
	html += '<label class="radio-inline">';
  	html += '<input type="radio" name="status" value="0" '+(status == 0 ? 'checked' : '')+'> 屏蔽';
	html += '</label>';
	html += '<label class="radio-inline">';
	html += '<input type="radio" name="status" value="1" '+(status == 1 ? 'checked' : '')+'> 正常';
	html += '</label>';
	html += '          </div>';
	
	
	
    html += '      </div>';
    html += '      <div class="modal-footer"><input type="hidden" name="id" value="'+id+'">';
    html += '        <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>';
    html += '        <button type="button" class="btn btn-primary js-album-submit">提交</button>';
    html += '      </div>';
    html += '    </div>';
	html += '   </form>';
    html += '  </div>';
    html += '</div>';
	$('body').append(html);
	$('#dialogAlbum').modal('show');
}

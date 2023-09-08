layui.use(['form', 'laydate', 'util','element'], function(){
    var form = layui.form;
    var layer = layui.layer;
    var util = layui.util;
    var element = layui.element;
    const download = new Download(element)
    const oldPath = new Stack();
    var allFile = []
    //打开
    $('body').on('click','.whole-click',async function(){
        let isdir = $(this).data("isdir")
        let idx = $(this).data("idx")
        let path = $(this).data("path")
        let name = $(this).data("name")
        if (!isdir){
            //查看图片
            let type = getFileType(name)
            if(type === "picture"){
                lookImg(layer,allFile,name,idx)
            }else if(type === "video"){
                lookVideo(layer,path,name)
                //|| type === "xls" || type === "doc" || type === "ppt" || type === "md"
            }else if(type === "txt" || type === "pdf" ){
                getView(layer,path,type)
            }else{
                //是文件
                layer.msg("不支持查看" , {icon: 2});
            }
            return false
        }
        oldPath.push(path)
        form.val('demo-val-filter',{path})
        allFile = await getAllFile(path)
        pushHtml(allFile)
        element.render('progress', 'demo-filter-progress');
        return false
    })

    //下载
    $('body').on('click','.download',function(){
        let path = $(this).data("path")
        let name = $(this).data("name")
        download.downloadFile(`/api/readFile?path=${path}`,name,this)
        return false
    })
    // 普通事件
    util.on('lay-on', {
        // 返回
        "get-retreat":async function(){
            //丢弃不用
            oldPath.pop()
            //需要返回的地址
            let path = oldPath.pop()
            if (path){
                form.val('demo-val-filter',{path})
                //当前地址重新入栈
                oldPath.push(path)
                allFile = await getAllFile(path)
                pushHtml(allFile)
            }else {
                form.val('demo-val-filter',{path:""})
                init()
            }
        },
        //前进
        "get-forward":async function(){
            var isvalid = form.validate('.input-path'); // 主动触发验证，v2.7.0 新增
            // 验证通过
            if(isvalid){
                let data = form.val('demo-val-filter');
                oldPath.push(data.path)
                allFile = await getAllFile(data.path)
                pushHtml(allFile)
            }
        },
        //搜索
        "i-search":async function(){
            let data = form.val('demo-val-filter'),
                search=data.search;
            if(search){
                //模糊查询
                let arr = allFile.filter(o => o.name.indexOf(search)>-1)
                pushHtml(arr)
                return
            }
            pushHtml(allFile)
        }
    });
});


//获取根目录
function init(){
    axios.get('/api/getLogicalDrives', {})
        .then(function (res) {
            if (res.data.code !== 0){
                return
            }
            pushHtml(res.data.data)
        })
        .catch(function (error) {
            console.log(error);
        });
}
init()

//渲染列表
function pushHtml(obj){
    $("#content").empty()
    for (let i in obj){
        obj[i]['icon'] = getFileType(obj[i].name)
        let downloadIcon = ` <div class="layui-col-xs1 layui-col-md1 download" data-path="${obj[i].path}" data-name="${obj[i].name}" style="text-align: right;">
                                        <i class="layui-icon layui-icon-download-circle list-icon" style="display: block;margin-top: 2px;"></i>
                                    </div>`,
            preview = `<i class="layui-icon layui-icon-extend layui-extend-${obj[i].isDir ? 'folder':obj[i].icon} list-icon"></i>`

        if (obj[i].icon === 'picture' ){
            preview = `<img class="preview-image" src="/api/readFile?path=${obj[i].path}" alt="${obj[i].name}">`
        }

        let h = `<div class="layui-panel whole-click" data-idx="${i}" data-path="${obj[i].path}" data-isdir="${obj[i].isDir}" data-name="${obj[i].name}">
                            <div class="layui-row">
                                <div class="layui-col-xs1 layui-col-md1" style="text-align: left">
                                    ${preview}
                                </div>
                                <div class="layui-col-xs10 layui-col-md10 ">
                                    <span class="list-span">${obj[i].name}</span>
                                </div>
                               ${obj[i].isDir ? '' : downloadIcon}
                            </div>
                        </div>`
        $("#content").append(h)
    }
}

//播放视频
function lookVideo(layer,path,name){
    // 后缀获取
    let suffix = '';
    try {
        const flieArr = name.split('.');
        suffix = flieArr[flieArr.length - 1];
    } catch (err) {
        suffix = '';
    }
    if(suffix === 'mp4' || suffix === 'webm'){
        layer.open({
            type: 1,
            title:name,
            area: ['100%','80%'], // 宽高
            closeBtn: 1,
            skin: 'class-layer-video-custom',
            content: '<div id="dplayer"></div>',
            success:function (layero, index){
                layer.iframeAuto(index);
                const dp = new DPlayer({
                    container: document.getElementById('dplayer'),
                    video: {
                        url: '/api/lookVideo?path='+path,
                    },
                });
            }
        });
    }else {
        layer.msg(`此【${suffix}】视频格式不支持播放`,{icon:2})
    }

}

//查看文件
function getView(layer,path,type){
    layer.open({
        type: 2,
        area: ['100%','90%'], // 宽高
        content: `/src/${type}.html`,
        fixed: false, // 不固定
        maxmin: true,
        shadeClose: true,
        success: function(layero, index, that){
            // 获取 iframe 的窗口对象
            let iframeWin =  window[layero.find('iframe')[0]['name']];
            iframeWin.initialization(path);
            layer.iframeAuto(index);
        }
    });
}

//查看图片
function lookImg(layer,allFile,name,idx){
    let ulEle = document.createElement("ul");
    let index = 0
    for(let i=0;i<allFile.length;i++){
        if(allFile[i].icon === 'picture'){
            let liEle = document.createElement("li");
            let img = new Image();
            img.src = "/api/readFile?path="+allFile[i].path;
            img.alt = name
            liEle.appendChild(img);
            ulEle.appendChild(liEle);
            if(idx > i){
                index++
            }
        }
    }

    const gallery = new Viewer(ulEle);
    gallery.view(index)
}

//获取列表
async function getAllFile(path){
    let data = []
    await axios.get('/api/getAllFile', {
        params:{path}
    })
    .then(function (res) {
        if (res.data.code !== 0){
            layer.msg(res.msg, {icon: 2});
            return
        }
        data = res.data.data;
    })
    .catch(function (error) {
        console.log(error);
    })
    return data
}


//文件类型
function getFileType(fileName) {
    // 后缀获取
    let suffix = '';
    // 获取类型结果
    let result = '';
    try {
        const flieArr = fileName.split('.');
        suffix = flieArr[flieArr.length - 1];
    } catch (err) {
        suffix = '';
    }
    // fileName无后缀返回 false
    if (!suffix) { return false; }
    suffix = suffix.toLocaleLowerCase();
    // 图片格式
    const imglist = ['png', 'jpg', 'jpeg', 'bmp', 'gif'];
    // 进行图片匹配
    result = imglist.find(item => item === suffix);
    if (result) {
        return 'picture';
    }
    // 匹配txt
    const txtlist = ['txt'];
    result = txtlist.find(item => item === suffix);
    if (result) {
        return 'txt';
    }
    // 匹配 excel
    const excelist = ['xls', 'xlsx'];
    result = excelist.find(item => item === suffix);
    if (result) {
        return 'xls';
    }
    // 匹配 word
    const wordlist = ['doc', 'docx'];
    result = wordlist.find(item => item === suffix);
    if (result) {
        return 'doc';
    }
    // 匹配 pdf
    const pdflist = ['pdf'];
    result = pdflist.find(item => item === suffix);
    if (result) {
        return 'pdf';
    }
    // 匹配 ppt
    const pptlist = ['ppt', 'pptx'];
    result = pptlist.find(item => item === suffix);
    if (result) {
        return 'ppt';
    }
    // 匹配 视频
    const videolist = ['mp4', 'm2v', 'mkv', 'rmvb', 'wmv', 'avi', 'flv', 'mov', 'm4v'];
    result = videolist.find(item => item === suffix);
    if (result) {
        return 'video';
    }
    // 匹配 音频
    const radiolist = ['mp3', 'wav', 'wmv'];
    result = radiolist.find(item => item === suffix);
    if (result) {
        return 'music';
    }
    //
    const md = ['msg', 'eml', 'md'];
    result = md.find(item => item === suffix);
    if (result) {
        return 'md';
    }

    // 其他 文件类型
    return 'file';
}

class Download {
    constructor(element)
    {
        //用于栈的索引和元素值的保存
        this.element = element
    }

    async downloadFile(fileUrl,fileName,el) {
        let whole = $(el).parent().parent(".whole-click"),
            idx = whole.data("idx");
        let h = `<div class="layui-progress progress-${idx}" lay-showpercent="true" lay-filter="filter-progress-${idx}">
                      <div class="layui-progress-bar" lay-percent="0%"></div>
                    </div>`
        whole.append(h)
        let data = await this.getData(fileUrl,idx);
        let blob = new Blob([data])
        this.saveFile(blob, fileName);
        $(`.progress-${idx}`).remove();
    }

    getData(fileUrl,idx) {
        let that = this
        return new Promise(resolve => {
            const xhr = new XMLHttpRequest();
            xhr.open('GET', fileUrl, true);
            //监听进度事件
            xhr.addEventListener(
                'progress',
                function (evt) {
                    if (evt.lengthComputable) {
                        let percentComplete = evt.loaded / evt.total;
                        // percentage是当前下载进度，可根据自己的需求自行处理
                        let percentage = percentComplete * 100;
                        that.element.progress('filter-progress-'+ idx, percentage + '%'); //执行进度条
                    }
                },
                false
            );
            xhr.responseType = 'arraybuffer';
            xhr.onload = () => {
                if (xhr.status === 200) {
                    resolve(xhr.response);
                }else {
                    $(`.progress-${idx}`).remove();
                    layer.msg("未知错误："+xhr.status,{icon:2})
                }
            };
            xhr.send();
        });
    }

    saveFile(blob, fileName) {
        // ie的下载
        if (window.navigator.msSaveOrOpenBlob) {
            navigator.msSaveBlob(blob, filename);
        } else {
            // 非ie的下载
            const link = document.createElement('a');
            const body = document.querySelector('body');

            link.href = window.URL.createObjectURL(blob);
            link.download = fileName;

            // fix Firefox
            link.style.display = 'none';
            body.appendChild(link);

            link.click();
            body.removeChild(link);

            window.URL.revokeObjectURL(link.href);
        }
    }
}





layui.config({
    // 配置 layui 第三方扩展组件存放的基础目录
    base: './plugins/layui_exts/'
}).extend({
    iconExtend: 'iconExtend/iconExtend'
}).use(['jquery', 'iconExtend'], function () {
    var $ = jquery = layui.jquery;
    var iconExtend = layui.iconExtend;
    // 加载Project对象 项目名称，配置
    var gf = iconExtend.loadProject('gf', {icon_class: 'a', style: {'color': 'cyan'}})
    // 图标放入位置的Dom对象
    var appDom = document.getElementById('app');
    // 通过Project对象 添加图标（最前） 返回IconDom对象
    var icon_dom = gf.prepend(appDom, 'layui-extend-cloud-server', {'color': 'blue'})
    // IconDom对象 仅提供基础的 style() fontSize() color() change() show() hide() remove() 方法 其他dom操作可以使用 $(icon_dom.dom)...
    //console.log($(icon_dom.dom))
    // 通过Project对象 添加图标（最后） 返回IconDom对象
    var icon_dom2 = gf.append(appDom)
    // 通过Project对象 添加图标（最后） 返回IconDom对象
    gf.append(appDom, 'notification')
    // 通过IconDom获取icon style
    //console.log(icon_dom, icon_dom.style('font-size'), icon_dom.fontSize(), icon_dom.color());
    icon_dom.on('click', function () {
        console.log(this);
    })
    setTimeout(function () {
        // 快捷设置字体大小 等价于 icon_dom.style('font-size','40px');
        icon_dom.fontSize('40px');
        // 快捷设置字体颜色
        icon_dom.color('red');
        // 改变图标 build-fill 等价于 layui-extend-build-fill
        icon_dom.change('build-fill');
        icon_dom2.change('cloud-sync');
    }, 1000)
    setTimeout(function () {
        // 设置图标样式
        icon_dom2.style({
            'font-size': '40px',
            'color': 'chocolate'
        });
    }, 2000)
    setTimeout(function () {
        icon_dom2.hide();
    }, 3000)
    setTimeout(function () {
        icon_dom2.show();
    }, 4000)
    setTimeout(function () {
        icon_dom2.remove();
    }, 5000)
});
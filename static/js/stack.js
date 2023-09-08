class Stack
{
    constructor()
    {
        //用于栈的索引和元素值的保存
        this.items = [];
    }

    //入栈
    push(element)
    {
        this.items.push(element);
    }

    //查看栈是否为空和它的大小
    isEmpty()
    {
        return this.items.length === 0;
    }

    size()
    {
        return this.items.length;
    }

    //弹出
    pop()
    {
        //查看栈是否为空
        if(this.isEmpty())
        {
            return undefined;
        }
        //此方法是删除最后一个元素并返回
        return this.items.pop();
    }

    //查看栈顶
    peek()
    {
        if(this.isEmpty())
        {
            return undefined;
        }
        return this.items[this.items.length -1];
    }

    //清空栈
    //将值复原为构造函数中的使用值即可
    clear()
    {
        this.items = [];
    }

}

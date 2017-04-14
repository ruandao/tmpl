一个简便的模版辅助函数

```
通常调用方法：
    s1, _ := template.ParseFiles("header.tmpl", "content.tmpl", "footer.tmpl")
    s1.ExecuteTemplate(os.Stdout, "header", nil)

问题在于这里需要罗列每个模版文件
使用了这个辅助函数后只需要设定相关目录就可以

    tmpl.AddDir("./templates")
	err := tmpl.Execute(os.Stdout, "content", nil)
	if err != nil {
		panic(err)
	}

使用案例，查看example 目录

经过检验，可以通过原生的｀ParseGlob｀函数达到相同的目的，对Glob少见多怪了，学习下
```
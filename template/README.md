- 每次检查bin目录下的_etc目录中的yaml文件,将其变动部分拷贝到etc目录中的yaml文件中,以下部分保持不动
    ```yaml
    cdn_url : http://localhost:9898/template
    source_path : ./
    ```

- 在bin目录下,选择自己当前系统版本的文件,执行./<文件名>就可以运行程序,程序启动后,在浏览器输入`http://localhost:9898`可以访问项目主页
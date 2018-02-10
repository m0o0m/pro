var es = require('event-stream');
var gulp = require('gulp');
var concat = require('gulp-concat');
var connect = require('gulp-connect');
var templateCache = require('gulp-angular-templatecache');
var ngAnnotate = require('gulp-ng-annotate');
var uglify = require('gulp-uglify');
var htmlmin = require('gulp-htmlmin');
var fs = require('fs');
var _ = require('lodash');

//第三方 js文件
var scripts = require('./app.scripts.json');

//代码
var adminSource = {
    js: {
        src: [
            // application config
            'app.config.js',
            'src/admin/app.service.config.js',
            // application bootstrap file
            'src/admin/app.js',
            'src/httpSvc.js',
            'src/filter.js',
            'src/directives.js',
            'src/resource.js',
            // module files
            'src/smart/**/module.js',
            'src/admin/**/module.js',

            // other js files [controllers, services, etc.]
            'src/smart/**/!(module)*.js',
            'src/admin/**/!(module)*.js'
        ],
        index: 'src/admin/index.html',
        views: 'src/admin/**/*.html',
        style: 'src/styles/**/*',
        staticViews: 'src/smart/**/*.html',
        plugin: 'smartadmin-plugin/**/*',
    },
    build: {
        script: "./admin",
        views: './admin/views',
        style: './admin'
    }
};

//
var agencySource = {
    js: {
        src: [
            // application config
            'app.config.js',
            'src/agency/app.service.config.js',
            // application bootstrap file
            'src/agency/app.js',
            'src/httpSvc.js',
            'src/filter.js',
            'src/directives.js',
            'src/resource.js',
            // module files
            'src/smart/**/module.js',
            'src/agency/**/module.js',


            // other js files [controllers, services, etc.]
            'src/smart/**/!(module)*.js',
            'src/agency/**/!(module)*.js'
        ],
        index: 'src/agency/index.html',
        views: 'src/agency/**/*.html',
        staticViews: 'src/smart/**/*.html',
        style: 'src/styles/**/*',
        plugin: 'smartadmin-plugin/**/*'
    },
    build: {
        script: "./agency",
        views: './agency/views',
        style: './agency'
    }
};


// 模版打包
gulp.task('agency:cache-templates', function () {
    var options = {
        removeComments: true,//清除HTML注释
        collapseWhitespace: true,//压缩HTML
        collapseBooleanAttributes: false,//省略布尔属性的值 <input checked="true"/> ==> <input />
        removeEmptyAttributes: true,//删除所有空格作属性值 <input id="" /> ==> <input />
        removeScriptTypeAttributes: true,//删除<script>的type="text/javascript"
        removeStyleLinkTypeAttributes: true,//删除<style>和<link>的type="text/css"
        minifyJS: true,//压缩页面JS
        minifyCSS: true//压缩页面CSS
    };
    return gulp.src([agencySource.js.staticViews,agencySource.js.views])
        .pipe(htmlmin(options))
        .pipe(templateCache('app.js',{
            root: 'views',
            module: 'app'
        }))
        .pipe(gulp.dest(agencySource.build.script));   
});

gulp.task('agency:build',['agency:cache-templates'], function () {
    agencySource.js.src.push(agencySource.build.script + '/app.js');
    return gulp.src(agencySource.js.src)
        .pipe(ngAnnotate())
        .pipe(uglify())
        .pipe(concat('app.js'))
        .pipe(gulp.dest(agencySource.build.script));
});


gulp.task('admin:cache-templates', function () {
    var options = {
        removeComments: true,//清除HTML注释
        collapseWhitespace: true,//压缩HTML
        collapseBooleanAttributes: false,//省略布尔属性的值 <input checked="true"/> ==> <input />
        removeEmptyAttributes: true,//删除所有空格作属性值 <input id="" /> ==> <input />
        removeScriptTypeAttributes: true,//删除<script>的type="text/javascript"
        removeStyleLinkTypeAttributes: true,//删除<style>和<link>的type="text/css"
        minifyJS: true,//压缩页面JS
        minifyCSS: true//压缩页面CSS
    };
    return gulp.src([adminSource.js.staticViews,adminSource.js.views])
        .pipe(htmlmin(options))
        .pipe(templateCache('app.js',{
            root: 'views',
            module: 'app'
        }))
        .pipe(gulp.dest(adminSource.build.script));
});
gulp.task('admin:build', ['admin:cache-templates'],function () {
    adminSource.js.src.push(adminSource.build.script + '/app.js');
    return es.merge(gulp.src(adminSource.js.src))
        .pipe(ngAnnotate())
        .pipe(uglify())
        .pipe(concat('app.js'))
        .pipe(gulp.dest(adminSource.build.script));
});

//js
gulp.task('agency:js', function () {
    //, agencyTemplateStream()
    return es.merge(gulp.src(agencySource.js.src))
        .pipe(concat('app.js'))
        .pipe(gulp.dest(agencySource.build.script));
});

gulp.task('admin:js', function () {
    //, adminTemplateStream()
    return es.merge(gulp.src(adminSource.js.src))
        .pipe(concat('app.js'))
        .pipe(gulp.dest(adminSource.build.script));
});


gulp.task('agency:style', function () {
    return gulp.src(agencySource.js.style)
        .on('error', handleError)
        .pipe(gulp.dest(agencySource.build.style));
});

gulp.task('admin:style', function () {
    return gulp.src(adminSource.js.style)
        .on('error', handleError)
        .pipe(gulp.dest(adminSource.build.style));
});

gulp.task('agency:views-index', function () {
    return gulp.src(agencySource.js.index)
    .on('error', handleError)
    .pipe(gulp.dest(agencySource.build.style));
});

gulp.task('agency:views', function () {
    err1 = gulp.src(agencySource.js.views)
        .on('error', handleError)
        .pipe(gulp.dest(agencySource.build.views));

    err2 = gulp.src(agencySource.js.staticViews)
        .on('error', handleError)
        .pipe(gulp.dest(agencySource.build.views));

    return err1 || err2 ;
});

gulp.task('admin:views-index', function () {
    return gulp.src(adminSource.js.index)
    .on('error', handleError)
    .pipe(gulp.dest(adminSource.build.style));
});

gulp.task('admin:views',['admin:views-index'], function () {
    err1 = gulp.src(adminSource.js.views)
        .on('error', handleError)
        .pipe(gulp.dest(adminSource.build.views));

    err2 = gulp.src(adminSource.js.staticViews)
        .on('error', handleError)
        .pipe(gulp.dest(adminSource.build.views));

    return err1 || err2;
});

gulp.task('agency:plugin', function () {
    return gulp.src(agencySource.js.plugin)
        .on('error', handleError)
        .pipe(gulp.dest(agencySource.build.views));
});

gulp.task('admin:plugin', function () {
    return gulp.src(adminSource.js.plugin)
        .on('error', handleError)
        .pipe(gulp.dest(adminSource.build.views));
});


//监听
gulp.task('agency:watch', function () {
    gulp.watch(agencySource.js.src, ['agency:js']);
    gulp.watch(agencySource.js.views, ['agency:js', 'agency:style', 'agency:views']);
});

gulp.task('admin:watch', function () {
    gulp.watch(adminSource.js.src, ['admin:js']);
    gulp.watch(adminSource.js.views, ['admin:js', 'admin:style', 'admin:views']);
});

gulp.task('agency:connect', function () {
    connect.server({
        root: './agency',
        port: 5000
    });
});

gulp.task('admin:connect', function () {
    connect.server({
        root: './admin',
        port: 3000
    });
});

//第三方
gulp.task('admin:vendor', function () {
    _.forIn(scripts.chunks, function (chunkScripts, chunkName) {
        var paths = [];
        chunkScripts.forEach(function (script) {
            var scriptFileName = scripts.paths[script];

            if (!fs.existsSync(__dirname + '/' + scriptFileName)) {

                throw console.error('Required path doesn\'t exist: ' + __dirname + '/' + scriptFileName, script)
            }
            paths.push(scriptFileName);
        });
        gulp.src(paths)
            .pipe(concat(chunkName + '.js'))
            .pipe(gulp.dest("./admin"))
    })

});

gulp.task('agency:vendor', function () {
    _.forIn(scripts.chunks, function (chunkScripts, chunkName) {
        var paths = [];
        chunkScripts.forEach(function (script) {
            var scriptFileName = scripts.paths[script];

            if (!fs.existsSync(__dirname + '/' + scriptFileName)) {

                throw console.error('Required path doesn\'t exist: ' + __dirname + '/' + scriptFileName, script)
            }
            paths.push(scriptFileName);
        });
        gulp.src(paths)
            .pipe(concat(chunkName + '.js'))
            .pipe(gulp.dest("./agency"));
    })

});

//
gulp.task('agency:prod', ['agency:vendor',  'agency:style', 'agency:views-index', 'agency:build','agency:connect']);
gulp.task('agency:dev', ['agency:vendor', 'agency:js', 'agency:style', 'agency:views', 'agency:watch', 'agency:connect']);

gulp.task('admin:prod', ['admin:vendor', 'admin:style','admin:views-index',  'admin:build','admin:connect']);
gulp.task('admin:dev', ['admin:vendor', 'admin:js', 'admin:style', 'admin:views', 'admin:watch', 'admin:connect']);
gulp.task('default', ['agency:dev']);

var swallowError = function (error) {
    console.log(error.toString());
    this.emit('end');
};

function handleError(err) {
    console.log(err.toString());
    this.emit('end');
}
const path = require('path')
const gulp = require('gulp')
const less = require('gulp-less')
const babel = require('gulp-babel')
const minifyJS = require('gulp-uglifyjs')
const minifyCSS = require('gulp-clean-css')
const minifyHTML = require('gulp-cleanhtml')
const replace = require('gulp-replace')
const concat = require('gulp-concat')
const embedTemplates = require('gulp-angular-embed-templates')
const autoPrefixer = require('gulp-autoprefixer')
const eslint = require('gulp-eslint')


gulp.task('default', ['watch'])

gulp.task('build', ['less', 'js', 'html'])
gulp.task('watch', ['build'], () => {
  gulp.watch('./client/src/**/*.js', ['js'])
  gulp.watch('./client/src/**/*.less', ['less'])
  gulp.watch('./client/**/*.html', ['html', 'js'])
})

gulp.task('less', () => {
  return gulp
    .src('./client/src/styles/styles.less')
    .pipe(less({
      paths: [path.join(__dirname, 'less', 'includes')]
    }))
    .pipe(autoPrefixer({
      browsers: ['last 2 versions']
    }))
    .pipe(minifyCSS())
    .pipe(gulp.dest('./client/build/styles/'))
})

gulp.task('js', () => {
  return gulp
    .src(['./client/src/scripts/app.js', './client/src/scripts/**/*.js'])
    .pipe(eslint('eslint.json'))
    .pipe(eslint.formatEach())
    .pipe(eslint.failOnError())
    .pipe(embedTemplates({
      minimize: {empty: true}
    }))
    .pipe(replace(/\>[\s]+\</g, '><'))
    .pipe(concat('app.js'))
    .pipe(babel({
      presets: ['es2015', 'es2016']
    }))
    .pipe(minifyJS())
    .pipe(gulp.dest('./client/build/scripts/'))
})

gulp.task('html', () => {
  return gulp
    .src('./client/src/index.html')
    .pipe(minifyHTML())
    .pipe(gulp.dest('./client/build/'))
})

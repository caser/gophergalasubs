var gulp = require('gulp');
var source = require('vinyl-source-stream'); // Used to stream bundle for further handling
var browserify = require('browserify');
var watchify = require('watchify');
var reactify = require('reactify'); 
var concat = require('gulp-concat');

gulp.task('browserify', function() {
  var bundler = browserify({
    entries: ['./app/main.js'], // Only need initial file, browserify finds the deps
    transform: [reactify], // We want to convert JSX to normal javascript
    debug: true, // Gives us sourcemapping
    cache: {}, packageCache: {}, fullPaths: true // Requirement of watchify
  });
  var watcher  = watchify(bundler);

  return watcher
  .on('update', function () { // When any files update
    var updateStart = Date.now();
    console.log('Updating!');
    watcher.bundle() // Create new bundle that uses the cache for high performance
    .pipe(source('main.js'))
    // This is where you add uglifying etc.
    .pipe(gulp.dest('./build/'));
    console.log('Updated!', (Date.now() - updateStart) + 'ms');
  })
  .bundle() // Create the initial bundle when starting the task
  .on('error', function(err){
    console.log(err)
  })
  .pipe(source('main.js'))
  .pipe(gulp.dest('./build/'));
});

// I added this so that you see how to run two watch tasks
gulp.task('css', function () {
  gulp.watch('./styles/**/*.css', function () {
    console.log("Copying styles")
    return gulp.src('styles/**/*.css')
    .pipe(concat('main.css'))
    .pipe(gulp.dest('build/'));
  });
});

// I added this so that you see how to run two watch tasks
gulp.task('images', function () {
  gulp.watch('./images/**/*', function () {
    console.log("Copying images")
    return gulp.src('./images/**/*')
    .pipe(gulp.dest('build/images/'));
  });
});

// Just running the two tasks
gulp.task('default', ['browserify', 'css', 'images']);

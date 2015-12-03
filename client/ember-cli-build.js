/* global require, module */
var EmberApp = require('ember-cli/lib/broccoli/ember-app');
var mergeTrees = require('broccoli-merge-trees');
var pickFiles = require('broccoli-static-compiler');
var extraAssets = pickFiles('bower_components/bootstrap-sass/assets/fonts/bootstrap', {
  srcDir: '/',
  files: ['**/*'],
  destDir: '/fonts/bootstrap'
});

module.exports = function(defaults) {
  var app = new EmberApp(defaults, {
  });

  app.import('bower_components/bootstrap-sass/assets/javascripts/bootstrap.js');

  return mergeTrees([app.toTree(), extraAssets]);
};

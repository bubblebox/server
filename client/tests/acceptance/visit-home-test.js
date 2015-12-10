import Ember from 'ember';
import { module, test } from 'qunit';
import startApp from 'firedragon/tests/helpers/start-app';

module('Acceptance | visit home', {
  beforeEach: function() {
    this.application = startApp();
  },

  afterEach: function() {
    Ember.run(this.application, 'destroy');
  }
});

test('visiting root', function(assert) {
  visit('/');

  andThen(function() {
    assert.equal(find('a.navbar-brand.active:contains(Fire Dragon)').length, 1);
  });
});

test('visiting root shows items', function (assert) {
  visit('/');
  //make sure we have an item
  andThen(function () {
    assert.equal(find('ul.items').length, 1);
  });
});
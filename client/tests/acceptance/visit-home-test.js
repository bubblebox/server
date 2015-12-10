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
    assert.equal(find('a.navbar-brand:contains(Fire Dragon)').length, 1);
  });
});

test('getting redirected to items', function(assert) {
  visit('/');

  andThen(function() {
    assert.equal(currentURL(), '/items');
  });
});

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
  const contact = server.create('contact');

  visit('/');

  andThen(function() {
    assert.equal(find('h2:contains(Firedragon)').length, 1);
  });
});

import Ember from 'ember';
import DS from 'ember-data';

export default DS.Model.extend({
  code: DS.attr('string'),
  payload: DS.attr('string'),
  type: DS.attr(),
  created: DS.attr(),
  views: DS.attr(),

  isURL: Ember.computed('type', function() {
    return this.get('type') === 0;
  }),

  isText: Ember.computed('type', function() {
    return this.get('type') === 1;
  }),

  payloadSummary: Ember.computed('payload', function() {
    var payload = this.get('payload');
    var shortPayload = Ember.$.trim(payload).substring(0, 10)
          .split(" ").slice(0, -1).join(" ") + "...";

    return shortPayload;
  }),
});

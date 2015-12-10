import Ember from 'ember';
import DS from 'ember-data';

export default DS.Model.extend({
  code: DS.attr('string'),
  content: DS.attr('string'),
  type: DS.attr(),
  created: DS.attr(),
  views: DS.attr(),

  isURL: Ember.computed('type', function() {
    return this.get('type') === 0;
  }),

  isText: Ember.computed('type', function() {
    return this.get('type') === 1;
  }),

  contentSummary: Ember.computed('content', function() {
    var content = this.get('content');
    var shortContent = Ember.$.trim(content).substring(0, 10)
          .split(" ").slice(0, -1).join(" ") + "...";

    return shortContent;
  }),
});

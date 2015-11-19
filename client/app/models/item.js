import DS from 'ember-data';

export default DS.Model.extend({
  code: DS.attr('string'),
  content: DS.attr('string'),
  type: DS.attr(),
});

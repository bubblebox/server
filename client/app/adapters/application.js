import DS from 'ember-data';

export default DS.RESTAdapter.extend({
  namespace: 'api/v1',
  host: 'http://api.127.0.0.1.xip.io:8042'
});

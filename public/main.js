var app = new Vue({
  el: '#app',
  data: {
    message: 'Hello Vue!',
    users: [],
    user: null,
  },
  mounted() {
    fetch('/users')
        .then(resp => resp.json())
        .then(users => this.users = users)
        .catch(err => console.error(err))
  }
});
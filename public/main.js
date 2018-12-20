new Vue({
  el: '#app',
  data: {
    message: 'Hello Vue!',
    users: [],
    user: null,
    passcode: '',
    touched: false,
    result: false
  },
  mounted() {
    fetch('/users')
        .then(resp => resp.json())
        .then(users => this.users = users)
        .catch(err => console.error(err))
  },
  methods: {
    validate() {
      fetch(`/users/${this.user.email}/validate`, {
        method: 'POST',
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({code: this.passcode})
      }).then(resp => resp.json())
          .then(resp => this.result = resp)
          .then(() => this.touched = true)
    },
    selectUser(email) {
      for (let i = 0; i < this.users.length; i++) {
        if (this.users[i].email === email) {
          this.user = this.users[i];
          return
        }
      }
    }
  }
});
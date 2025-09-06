import { defineStore } from 'pinia';

interface AuthUserInfo {
  nickname: string;
  id: number;
  avatar: string;
  permission: string[];
}

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null as AuthUserInfo | null,
    token: ''
  }),
  getters: {
    currentToken: (state) => {
      let tk: string | null = state.token;
      if (!tk || tk.length == 0) {
        tk = window.localStorage.getItem('YEX_AUTH');
      }
      if (!tk || tk.length == 0) {
        return '';
      }
      state.token = tk;
      return tk;
    }
  },
  actions: {
    setUser(user: any) {
      this.user = user
    },
    setToken(token: string) {
      this.token = token;
      if (token.length > 0) {
        window.localStorage.setItem('YEX_AUTH', token);
      } else {
        window.localStorage.removeItem('YEX_AUTH');
      }
    },
    loginOut () {
      this.user = null;
      this.token = '';
      window.localStorage.removeItem('YEX_AUTH');
    },
  }
})
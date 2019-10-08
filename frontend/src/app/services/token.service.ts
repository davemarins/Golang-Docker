import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class TokenService {

  constructor() { }

  handle(token) {
    this.set(token);
  }

  set(token) {
    localStorage.setItem('access_token', token.Token);
    localStorage.setItem('access_token_expire', token.Expires);
    localStorage.setItem('access_token_refresh', token.Refreshes);
  }

  getRefreshDate() {
    return localStorage.getItem('access_token_refresh');
  }

  getExpireDate() {
    return localStorage.getItem('access_token_expire');
  }

  get() {
    return localStorage.getItem('access_token');
  }

  remove() {
    localStorage.removeItem('access_token');
    localStorage.removeItem('access_token_expire');
    localStorage.removeItem('access_token_refresh');
  }

  isValid() {
    var expires_in = this.getExpireDate();
    var now = Math.floor(new Date().getTime() / 1000);
    if(expires_in == "" || expires_in == null || parseInt(this.getExpireDate()) <= now) {
      this.remove();
      return false;
    } else {
      return true;
    }
  }

  payload(token) {
    const payload = token.split('.')[1];
    return this.decode(payload);
  }

  decode(payload) {
    return JSON.parse(atob(payload));
  }

  loggedIn() {
    return this.isValid();
  }

}
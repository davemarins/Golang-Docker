import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class TokenService {

  private iss = {
    login: 'http://localhost:8000/public/user/login/'
  }

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
    const token = this.get();
    if(token) {
      const payload = this.payload(token);
      if(payload) {
        return Object.values(this.iss).indexOf(payload.iss) > -1 ? true: false;
      }
    }
    return false;
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
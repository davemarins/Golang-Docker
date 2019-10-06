import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class WebService {

  private public_url = 'http://localhost:8080/public/';
  private auth_url = 'http://localhost:8080/';

  constructor(private http: HttpClient) { }

  login(data) {
    return this.http.post(`${this.public_url}user/login/`, data);
  }

  refresh() {
    return this.http.post(`${this.auth_url}user/refresh/`, null);
  }

}

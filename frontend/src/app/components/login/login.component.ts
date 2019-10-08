import { Component, OnInit } from '@angular/core';
import { WebService } from 'src/app/services/web.service';
import { TokenService } from '../../services/token.service';
import { Router } from '@angular/router';
import { AuthService } from '../../services/auth.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {

  public form = {
    email: null,
    password: null,
    captcha: null
  };

  public email: string;
  public password: string;
  public captcha: string;
  public error = null;
  public success = null;

  constructor(
    private web: WebService,
    private token: TokenService,
    private router: Router,
    private auth: AuthService) { }

  ngOnInit() {
  }

  handlingResponse(data) {
    this.token.handle(data);
    this.auth.changeAuthStatus(true);
    this.router.navigateByUrl('/home');
  }

  handlingError(error) {
    this.error = error.error.error;
  }

  submitLogin() {
    this.form.email = this.email;
    this.form.password = this.password;
    this.web.login(this.form).subscribe(
      data => this.handlingResponse(data),
      error => this.handlingError(error)
    );
  }

}

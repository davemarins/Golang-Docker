import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { LoginComponent } from './components/login/login.component';
import { HomeComponent } from './components/home/home.component';

import { BeforeLoginService } from './services/before-login.service';
import { AfterLoginService } from './services/after-login.service';
import { ArticlesComponent } from './components/articles/articles.component';
import { SubscribersComponent } from './components/subscribers/subscribers.component';
import { MailboxComponent } from './components/mailbox/mailbox.component';

const routes: Routes = [
  {
    path: 'login',
    component: LoginComponent,
    canActivate: [BeforeLoginService]
  },
  {
    path: 'home',
    component: HomeComponent,
    canActivate: [AfterLoginService]
  },
  {
    path: 'articles',
    component: ArticlesComponent,
    canActivate: [AfterLoginService]
  },
  {
    path: 'subscribers',
    component: SubscribersComponent,
    canActivate: [AfterLoginService]
  },
  {
    path: 'mailbox',
    component: MailboxComponent,
    canActivate: [AfterLoginService]
  }
  // other views to be added
  // TODO:
  // 1) digital signature
  // 2) electronic invoice
  // 3) certified mailbox
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }

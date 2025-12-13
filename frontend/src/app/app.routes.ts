import { Routes } from '@angular/router';
import { DashboardComponent } from './pages/dashboard/dashboard.component';
import {BookingsComponent} from "./pages/bookings/bookings";
import {PropertiesComponent} from "./pages/properties/properties.component";
import {PaymentsComponent} from "./pages/payments/payments.component";
import {CustomersComponent} from "./pages/customers/customers.component";
import {UsersComponent} from "./pages/users/users.component";
import {DashboardCustomerComponent} from "./pages/dashboard-customer/dashboard-customer";

export const routes: Routes = [
  {
    path: '',
    redirectTo: 'dashboard',
    pathMatch: 'full'
  },
  {
    path: 'dashboard',
    component: DashboardComponent
  },
  {
    path: 'dashboard-customer',
    component: DashboardCustomerComponent
  },
  {
    path: 'bookings',
    component: BookingsComponent
  },
  {
    path: 'properties',
    component: PropertiesComponent
  },
  {
    path: 'payments',
    component: PaymentsComponent
  },
  {
    path: 'customers',
    component: CustomersComponent
  },
  {
    path: 'users',
    component: UsersComponent
  }
];

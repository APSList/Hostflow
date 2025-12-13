import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import {Button} from "primeng/button";
import {HeaderComponent} from "./components/header/header.component";
import {DashboardComponent} from "./pages/dashboard/dashboard.component";

@Component({
  selector: 'app-root',
  imports: [RouterOutlet, Button, HeaderComponent, DashboardComponent],
  templateUrl: './app.component.html',
  standalone: true,
  styleUrl: './app.component.scss'
})
export class AppComponent {
  title = 'Hostflow';
}

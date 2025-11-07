import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import {Button} from "primeng/button";
import {HeaderComponent} from "./components/header/header.component";

@Component({
  selector: 'app-root',
  imports: [RouterOutlet, Button, HeaderComponent],
  templateUrl: './app.component.html',
  standalone: true,
  styleUrl: './app.component.scss'
})
export class AppComponent {
  title = 'Hostflow';
}

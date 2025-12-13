import { Component } from '@angular/core';
import {PropertyListComponent} from "../../components/property-list-component/property-list-component";

@Component({
  selector: 'app-properties',
  standalone: true,
  imports: [
    PropertyListComponent
  ],
  templateUrl: './properties.component.html',
  styleUrl: './properties.component.css',
})
export class PropertiesComponent {

}

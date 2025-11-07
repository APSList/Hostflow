import {Component, OnInit} from '@angular/core';
import {MenuItem} from "primeng/api";
import {Toast} from "primeng/toast";
import {Menubar} from "primeng/menubar";

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrl: './header.component.scss',
  imports: [
    Toast,
    Menubar
  ],
  standalone: true
})
export class HeaderComponent implements OnInit{
  items: MenuItem[] | undefined;

  ngOnInit() {
    this.items = [
      {
        label: 'File',
        icon: 'pi pi-file',
        items: [
          {
            label: 'New',
            icon: 'pi pi-plus',
            command: () => {
            }
          },
          {
            label: 'Print',
            icon: 'pi pi-print',
            command: () => {
            }
          }
        ]
      },
      {
        label: 'Search',
        icon: 'pi pi-search',
        command: () => {
        }
      },
      {
        separator: true
      },
      {
        label: 'Sync',
        icon: 'pi pi-cloud',
        items: [
          {
            label: 'Import',
            icon: 'pi pi-cloud-download',
            command: () => {
            }
          },
          {
            label: 'Export',
            icon: 'pi pi-cloud-upload',
            command: () => {
            }
          }
        ]
      }
    ];
  }
}

import { Component, signal, computed } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Router } from '@angular/router';

import { TableModule } from 'primeng/table';
import { ButtonModule } from 'primeng/button';
import { InputTextModule } from 'primeng/inputtext';
import { Select } from 'primeng/select'; // ✅ PrimeNG 21 standalone Select

type PropertyStatus = 'Active' | 'Inactive' | 'Draft';

interface Property {
  id: number;
  name: string;
  country: string;
  status: PropertyStatus;
  maxGuests: number;
}

@Component({
  selector: 'app-properties-list',
  standalone: true,
  imports: [
    CommonModule,
    TableModule,
    ButtonModule,
    InputTextModule,
    Select // ✅ NAMSTO SelectModule
  ],
  templateUrl: './property-list-component.html'
})
export class PropertyListComponent {

  /* MOCK DATA */
  private allProperties = signal<Property[]>([
    {
      id: 1,
      name: 'Lake House',
      country: 'Slovenia',
      status: 'Active',
      maxGuests: 6
    },
    {
      id: 2,
      name: 'City Apartment',
      country: 'Austria',
      status: 'Inactive',
      maxGuests: 4
    },
    {
      id: 3,
      name: 'Mountain Cabin',
      country: 'Italy',
      status: 'Draft',
      maxGuests: 8
    }
  ]);

  /* FILTER STATE */
  nameFilter = signal('');
  statusFilter = signal<PropertyStatus | null>(null);

  statusOptions = [
    { label: 'Active', value: 'Active' },
    { label: 'Inactive', value: 'Inactive' },
    { label: 'Draft', value: 'Draft' }
  ];

  /* FILTERED RESULT */
  filteredProperties = computed(() =>
    this.allProperties().filter(p => {
      const nameMatch =
        !this.nameFilter() ||
        p.name.toLowerCase().includes(this.nameFilter().toLowerCase());

      const statusMatch =
        !this.statusFilter() ||
        p.status === this.statusFilter();

      return nameMatch && statusMatch;
    })
  );

  constructor(private router: Router) {}

  open(id: number) {
    this.router.navigate(['/properties', id]);
  }

  create() {
    this.router.navigate(['/properties', 'new']);
  }

  delete(id: number) {
    if (!confirm('Delete property?')) return;

    this.allProperties.update(list =>
      list.filter(p => p.id !== id)
    );
  }
}

import { Component, OnInit,ChangeDetectorRef } from '@angular/core';
import {SpinnerServiceService} from '../../services/spinner-service.service';

@Component({
  selector: 'app-spinner',
  templateUrl: './spinner.component.html',
  styleUrls: ['./spinner.component.css']
})
export class SpinnerComponent implements OnInit {
  showSpinner = false;
  constructor(private spinnerService: SpinnerServiceService,private cdRef:ChangeDetectorRef) { }


  ngOnInit(): void{
    this.spinnerService.getSpinnerObserver().subscribe(status=>{
      this.showSpinner=status==='start';/*Solo se mostrara cuando este el start */
      this.cdRef.detectChanges();/*En algunos componenetes no detecta los cambios */
    });
  }

}

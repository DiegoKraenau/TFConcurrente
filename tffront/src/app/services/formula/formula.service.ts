import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class FormulaService {

  private baseURL:string='http://localhost:9000';
  private extendURL:string='/regresionLinealConQuery?mathScore='
  private URL=this.baseURL+this.extendURL

  constructor(private http:HttpClient) { }

  
  doFormula(numero:Number): Observable<any>{

    return this.http.get(`${this.URL+numero}`);
  }
  

}

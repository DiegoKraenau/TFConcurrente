
import { Component, OnInit } from '@angular/core';
import {Variable} from '../../entities/Variable';
import {Formula} from '../../entities/Formula';
import {FormulaService} from '../../services/formula/formula.service';

@Component({
  selector: 'app-main',
  templateUrl: './main.component.html',
  styleUrls: ['./main.component.css']
})
export class MainComponent implements OnInit {


  variable: Variable = new Variable();
  formula: Formula=new Formula();
  procesado: boolean;

  constructor(private formulaService: FormulaService) { 
    this.variable.Mathscore=0
    this.procesado=false
  }

  ngOnInit(): void {

  }

  doFormula(){
    console.log(this.variable.Mathscore)
    
    this.formulaService.doFormula(this.variable.Mathscore).subscribe(
      response=>{ 
        console.log(response)
        this.procesado=true
        this.formula=response
      }
    )



  }

}

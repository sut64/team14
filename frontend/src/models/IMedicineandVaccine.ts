import { AgeInterface } from "./IAge";
import { CategoryInterface } from "./ICategory";
import { ContagiosInterface } from "./IContagios";
import { DosageFormInterface } from "./IDosageForm";

export interface MedicineandVaccineInterface {
	ID: number,
	RegNo: String,
	Name: String,
	Date: Date ,
	MinAge: number,
	MaxAge: number,
	DosageFormID: number,
	DosageForm:   DosageFormInterface, 
	ContagiosID: number,
	Contagios:  ContagiosInterface, 
	AgeID: number,
	Age:   AgeInterface, 
	CategoryID: number,
	Category:   CategoryInterface, 
   }
   

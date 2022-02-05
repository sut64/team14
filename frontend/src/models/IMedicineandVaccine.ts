import { AgeInterface } from "./IAge";
import { CategoryInterface } from "./ICategory";
import { ContagiousInterface } from "./IContagious";
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
	ContagiousID: number,
	Contagious:  ContagiousInterface, 
	AgeID: number,
	Age:   AgeInterface, 
	CategoryID: number,
	Category:   CategoryInterface, 
   }
   

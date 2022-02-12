import { OfficersInterface } from "./IOfficer";
import { ContagiousInterface } from "./IContagious";
import { SpecialistsInterface } from "./ISpecialist";

export interface PreventionsInterface {
    ID: number,
    OfficerID: number,
    Officer: OfficersInterface,
    ContagiousID: number,
    Contagious: ContagiousInterface,
    Disease: string,
    SpecialistID: number,
    Specialist: SpecialistsInterface,
    Protection: string,
    Age: number,
    Date: Date,
}

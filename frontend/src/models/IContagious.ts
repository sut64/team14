import { GermInterface } from "./IGerm";
import { CatchingTypeInterface } from "./ICatchingType";
import { RiskGroupTypeInterface } from "./IRiskGroupType";

export interface ContagiousInterface {
    ID: number,
    Name: string,
    Symptom: string,
    Incubation: number,
    Date: Date,
    GermID: number,
    Germ: GermInterface,
    CatchingTypeID: number
    CatchingType: CatchingTypeInterface
    RiskGroupTypeID: number
    RiskGroupType: RiskGroupTypeInterface
}
   

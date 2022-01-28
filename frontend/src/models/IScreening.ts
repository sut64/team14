import { PatientInterface } from "../models/IPatient";
import { RoomInterface } from "./IRoom";
import { SymptomInterface } from "../models/ISymptom";
import { OfficerInterface } from "../models/IOfficer";

export interface ScreeningInterface{
    ID: number,
    Time: Date,
    SymptomID: number,
    Symptom: SymptomInterface,
    PatientID: number,
    Patient: PatientInterface,
    RoomID: number,
    Room: RoomInterface,
    OfficerID: number,
    Officer: OfficerInterface,
}

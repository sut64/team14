import { PatientsInterface } from "../models/IPatient";
import { RoomInterface } from "./IRoom";
import { SymptomInterface } from "../models/ISymptom";
import { OfficersInterface } from "../models/IOfficer";

export interface ScreeningInterface{
    ID: number,
    Time: Date,
    SymptomID: number,
    Symptom: SymptomInterface,
    PatientID: number,
    Patient: PatientsInterface,
    RoomID: number,
    Room: RoomInterface,
    OfficerID: number,
    Officer: OfficersInterface,
    BloodPressure: number,
    CongenitalDisease: string,
}

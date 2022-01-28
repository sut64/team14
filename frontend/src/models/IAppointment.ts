import { OfficersInterface } from "./IOfficer";
import { SpecialistsInterface } from "./ISpecialist";
import { PatientsInterface } from "./IPatient";

export interface AppointmentInterface {
    ID: number,
    Note: string,
    Number: number,
    AppointDate: Date,
    IssueDate: Date,

    OfficerID: number,
    Officer: OfficersInterface,

    SpecialistID: number,
    Specialist: SpecialistsInterface,

    PatientID: number,
    Patient: PatientsInterface,

}
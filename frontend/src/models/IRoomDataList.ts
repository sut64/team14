import { OfficersInterface } from "./IOfficer";
import { PatientsInterface } from "./IPatient";
import { SpecialistsInterface } from "./ISpecialist";
import { RoomDetailsInterface } from "./IRoomDetail";

export interface RoomDataListsInterface {
  ID: string,
  OfficerID: number,
  Officer: OfficersInterface,
  
  RoomID: number,
  Room:  RoomDetailsInterface,

  PatientID: number,
  Patient: PatientsInterface,

  SpecialistID: number,
  Specialist:SpecialistsInterface,
  
  Day: number,
  Note: string,
  EnterRoomTime: Date,
}
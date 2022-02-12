import { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import {
  makeStyles,
  Theme,
  createStyles,
  //alpha,
} from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import FormControl from "@material-ui/core/FormControl";
import Container from "@material-ui/core/Container";
import Paper from "@material-ui/core/Paper";
import Grid from "@material-ui/core/Grid";
import Box from "@material-ui/core/Box";
import Typography from "@material-ui/core/Typography";
import Divider from "@material-ui/core/Divider";
import Snackbar from "@material-ui/core/Snackbar";
import Select from "@material-ui/core/Select";
import MuiAlert, { AlertProps } from "@material-ui/lab/Alert";
import TextField from '@material-ui/core/TextField';
import { OfficersInterface } from "../models/IOfficer";

import { PatientsInterface } from "../models/IPatient";
import { RoomInterface } from "../models/IRoom";
import { SymptomInterface } from "../models/ISymptom";
import { ScreeningInterface } from "../models/IScreening";

import {
  MuiPickersUtilsProvider,
  KeyboardDateTimePicker,
} from "@material-ui/pickers";
import DateFnsUtils from "@date-io/date-fns";
import React from "react";

const Alert = (props: AlertProps) => {
  return <MuiAlert elevation={6} variant="filled" {...props} />;
};

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      flexGrow: 1,
    },
    container: {
      marginTop: theme.spacing(2),
    },
    paper: {
      padding: theme.spacing(2),
      color: theme.palette.text.secondary,
    },
  })
);

function ScreeningCreate() {
  const classes = useStyles();
  const [time, setTime] = useState<Date | null>(new Date());
  const [officers, setOfficer] = useState<OfficersInterface[]>([]);

const [patient, setPatient] = useState<PatientsInterface[]>([]);
const [room, setRoom] = useState<RoomInterface[]>([]);
const [symptom, setSymptom] = useState<SymptomInterface[]>([]);
const [screening, setScreening] = useState<Partial<ScreeningInterface>>(
    {}
  );
  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(false);
  const [errormassage, setErrorMassage] = useState("");

  const apiUrl = "http://localhost:8080";
  const requestOptions = {
    method: "GET",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
    },
  };

  const handleClose = (event?: React.SyntheticEvent, reason?: string) => {
    if (reason === "clickaway") {
      return;
    }
    setSuccess(false);
    setError(false);
  };

  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>
  ) => {
    const name = event.target.name as keyof typeof screening;
    setScreening({
      ...screening,
      [name]: event.target.value,
    });
  };

  const handleTimeChange = (date: Date | null) => {
    console.log(date);
    setTime(date);
  };
 
  

  const getOfficer = async () => {
    let uid = localStorage.getItem("uid")
    fetch(`${apiUrl}/officer/${uid}`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        screening.OfficerID = res.data.ID
        if (res.data) {
          setOfficer(res.data);
        } else {
          console.log("else");
        }
      });
  };


  const getRoom = async () => {
    fetch(`${apiUrl}/rooms`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setRoom(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getSymptom = async () => {
    fetch(`${apiUrl}/symptoms`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setSymptom(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getPatient = async () => {
    fetch(`${apiUrl}/patients`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setPatient(res.data);
        } else {
          console.log("else");
        }
      });
  };

  useEffect(() => {
    getOfficer();
    getPatient();
    getSymptom();
    getRoom();
  }, []);

  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };
 
  function submit() {
    let data = {
      OfficerID: convertType(screening.OfficerID),
      PatientID: convertType(screening.PatientID),
      RoomID: convertType(screening.RoomID),
      SymptomID: convertType(screening.SymptomID),
      Time: time,
      BloodPressure:  convertType(screening.BloodPressure),
      CongenitalDisease:  screening.CongenitalDisease
    }; 
    console.log(data)
    
    const requestOptionsPost = {
      method: "POST",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    };
        fetch(`${apiUrl}/screenings`, requestOptionsPost)
          .then((response) => response.json())
          .then((res) => {
            if (res.data) {
              setSuccess(true);
	      setErrorMassage("");
            } else {
              setError(true);
	      setErrorMassage(res.error);
            }
          });
    }

  return (
    <Container className={classes.container} maxWidth="md">
      <Snackbar open={success} autoHideDuration={6000} onClose={handleClose}>
        <Alert onClose={handleClose} severity="success">
          บันทึกข้อมูลสำเร็จ
        </Alert>
      </Snackbar>
      <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
        <Alert onClose={handleClose} severity="error">
          บันทึกข้อมูลไม่สำเร็จ: {errormassage}
        </Alert>
      </Snackbar>
      <Paper className={classes.paper}>
        <Box display="flex">
          <Box flexGrow={1}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              บันทึกข้อมูลผู้ป่วย
            </Typography>
          </Box>
        </Box>
        <Divider />
        <Grid container spacing={3} className={classes.root}>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>ผู้ป่วย</p>
              <Select
                native
                value={screening.PatientID}
                onChange={handleChange}
                inputProps={{
                  name: "PatientID",
                }}
              >
                <option aria-label="None" value="">
                  เลือกผู้ป่วย
                </option>
                {patient.map((item: PatientsInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Name}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>อาการ</p>
              <Select
                native
                value={screening.SymptomID}
                onChange={handleChange}
                inputProps={{
                  name: "SymptomID",
                }}
              >
                <option aria-label="None" value="">
                  เลือกอาการ
                </option>
                {symptom.map((item: SymptomInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.State}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>ห้องพักผู้ป่วย</p>
              <Select
                native
                value={screening.RoomID}
                onChange={handleChange}
                inputProps={{
                  name: "RoomID",
                }}
              >
                <option aria-label="None" value="">
                  เลือกห้องพักผู้ป่วย
                </option>
                {room.map((item: RoomInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.RoomNumber}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
	  <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>ความดัน</p>
              <TextField
                name="BloodPressure"
                variant="outlined"
                type="string"
                size="medium"
                placeholder="โปรดใส่ความดันเลือด"
                value={screening.BloodPressure}
                onChange={handleChange}
              />
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>โรคประจำตัว</p>
              <TextField
                name="CongenitalDisease"
                variant="outlined"
                type="string"
                size="medium"
                placeholder="โปรดระบุโรคประจำตัว"
                value={screening.CongenitalDisease}
                onChange={handleChange}
              />
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>วันที่และเวลา</p>
              <MuiPickersUtilsProvider utils={DateFnsUtils}>
                <KeyboardDateTimePicker
                  name="Time"
                  value={time}
                  onChange={handleTimeChange}
                  label="กรุณาเลือกวันที่และเวลา"
                  minDate={new Date("2018-01-01T00:00")}
                  format="yyyy/MM/dd hh:mm a"
                />
              </MuiPickersUtilsProvider>
            </FormControl>
          </Grid>
          <Grid item xs={12}>
            <Button
              component={RouterLink}
              to="/screenings"
              variant="contained"
            >
              กลับ
            </Button>
            <Button
              style={{ float: "right" }}
              variant="contained"
              onClick={submit}
              color="primary"
            >
              บันทึก
            </Button>
          </Grid>
        </Grid>
      </Paper>
    </Container>
  );
}

export default ScreeningCreate;






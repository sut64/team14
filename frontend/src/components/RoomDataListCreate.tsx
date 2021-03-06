import React, { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import { makeStyles, Theme, createStyles } from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import FormControl from "@material-ui/core/FormControl";
import Container from "@material-ui/core/Container";
import Paper from "@material-ui/core/Paper";
import Grid from "@material-ui/core/Grid";
import Box from "@material-ui/core/Box";
import Typography from "@material-ui/core/Typography";
import Divider from "@material-ui/core/Divider";
import Snackbar from "@material-ui/core/Snackbar";
import MuiAlert, { AlertProps } from "@material-ui/lab/Alert";
import { Select } from "@material-ui/core";
import TextField from '@material-ui/core/TextField';
import DateFnsUtils from '@date-io/date-fns';
import { KeyboardDateTimePicker, MuiPickersUtilsProvider } from "@material-ui/pickers";

import { OfficersInterface } from "../models/IOfficer";
import { PatientsInterface } from "../models/IPatient";
import { SpecialistsInterface } from "../models/ISpecialist";
import { RoomDetailsInterface } from "../models/IRoomDetail";
import { RoomDataListsInterface } from "../models/IRoomDataList";



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

function RoomDataListCreate() {
  const classes = useStyles();
  const [selectedDate, setSelectedDate] = useState<Date | null>(new Date());
  
  const [Officer, setOfficers] = useState<OfficersInterface>();
  const [Patient, setPatients] = useState<PatientsInterface[]>([]);
  const [Specialist, setSpecialists] = useState<SpecialistsInterface[]>([]);
  const [RoomDetail, setRoomDetails] = useState<RoomDetailsInterface[]>([]);
  const [RoomDataList, setRoomDataList] = useState<Partial<RoomDataListsInterface>>( {});
 

  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(false);
  const [errorMessage, setErrorMessage] = useState("");

  const apiUrl = "http://localhost:8080";
  const requestOptions = {
    method: "GET",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json", },
  };

  

  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>
  ) => {
    const name = event.target.name as keyof typeof RoomDataList;
    setRoomDataList({
      ...RoomDataList,
      [name]: event.target.value,
    });
  };

  const handleClose = (event?: React.SyntheticEvent, reason?: string) => {
    if (reason === "clickaway") {
      return;
    }
    setSuccess(false);
    setError(false);
  };

  const handleInputChange = (
    event: React.ChangeEvent<{ id?: string; value: any }>
  ) => {
    const id = event.target.id as keyof typeof RoomDataList;
    const { value } = event.target;
    setRoomDataList({ ...RoomDataList, [id]: value });
  };

  const handleDateChange = (date: Date | null) => {
    console.log(date);
    setSelectedDate(date);
  };

  const getOfficer = async () => {
    let uid = localStorage.getItem("uid");
    fetch(`${apiUrl}/officer/ ${uid}`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        RoomDataList.OfficerID = res.data.ID
        if (res.data) {
          setOfficers(res.data);
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
          setPatients(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getSpecialist = async () => {
    fetch(`${apiUrl}/specialists`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setSpecialists(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getRoomDetail = async () => {
    fetch(`${apiUrl}/room_details`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setRoomDetails(res.data);
        } else {
          console.log("else");
        }
      });
  };

  

  useEffect(() => {
    getOfficer();
    getPatient();
    getSpecialist();
    getRoomDetail();
   
  }, []);

  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  function submit() {
    let data = {
      OfficerID: convertType(RoomDataList.OfficerID),
      PatientID: convertType(RoomDataList.PatientID),
      SpecialistID: convertType(RoomDataList.SpecialistID),
      RoomID: convertType(RoomDataList.RoomID),
      Day: typeof RoomDataList.Day === "string" ? parseInt(RoomDataList.Day) : 0 ,
      Note: RoomDataList.Note ,
      EnterRoomTime: selectedDate,
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

    fetch(`${apiUrl}/room_data_lists`, requestOptionsPost)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          console.log("???????????????????????????")
          setSuccess(true);
          setErrorMessage("");
        } else {
          console.log("????????????????????????????????????",res)
          setError(true);
          setErrorMessage(res.error);
          
        }
      });
  }
  return (
    <Container className={classes.container} maxWidth="md">
      <Snackbar open={success} autoHideDuration={6000} onClose={handleClose}>
        <Alert onClose={handleClose} severity="success">
          ??????????????????????????????????????????????????????????????????????????????
        </Alert>
      </Snackbar>
      <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
        <Alert onClose={handleClose} severity="error">
          ??????????????????????????????????????????????????????????????? {errorMessage}
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
              ????????????????????????????????????????????????????????????
            </Typography>
          </Box>
        </Box>
        <Divider />
        <Grid container spacing={3} className={classes.root}>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>Officer</p>
              <Select
                native
                value={RoomDataList.OfficerID}
                onChange={handleChange}
                disabled
                inputProps={{
                  name: "OfficerID",
                }}
              >
                <option aria-label="None" value="">
                 Officer
                </option>
                <option value={Officer?.ID} key={Officer?.ID}>
                  {Officer?.Name}
                </option>

              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>RoomDetail</p>
              <Select
                native
                value={RoomDataList.RoomID}
                onChange={handleChange}
                inputProps={{
                  name: "RoomID",
                }}
              >
                <option aria-label="None" value="">
                  Plase select Room
                </option>

                {RoomDetail.map((item: RoomDetailsInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Name}
                  </option>
                ))}

                
              </Select>
            </FormControl>
          </Grid>

        
          <Grid item xs={6}>
          <FormControl fullWidth variant="outlined">
              <p>Patient</p>
              <Select
                native
                value={RoomDataList.PatientID}
                onChange={handleChange}
              
                inputProps={{
                  name: "PatientID",
                }}
              >
                <option aria-label="None" value="">
                    Patient
                </option>
                {Patient.map((item: PatientsInterface) => (
                <option value={item?.ID} key={item?.ID}>
                  {item?.Name}
                </option>
                ))}
                
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>Specialist</p>
              <Select
                native
                value={RoomDataList.SpecialistID}
                onChange={handleChange}
                inputProps={{
                  name: "SpecialistID",
                }}
              >
                <option aria-label="None" value="">
                  Please select Specialist
                </option>
                {Specialist.map((item: SpecialistsInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Name}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>

        
          
         <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>Day (amount)</p>
              <TextField
                id="Day"
                variant="outlined"
                type="integer"
                size="medium"
                placeholder="Please insert amount of day "
                value={RoomDataList.Day || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>
          
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>Note</p>
              <TextField
                id="Note"
                variant="outlined"
                type="string"
                size="medium"
                placeholder="Please insert Note"
                value={RoomDataList.Note || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>


          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>Enter Room TIME</p>
              <MuiPickersUtilsProvider utils={DateFnsUtils}>
                <KeyboardDateTimePicker
                  name="enterroomtime"
                  value={selectedDate}
                  onChange={handleDateChange}
                  
                  
                  label="Please select Enter time"
                  minDate={new Date("2018-01-01T00:00")}
                  format="yyyy/MM/dd hh:mm a"
                  

                />
              </MuiPickersUtilsProvider>
            </FormControl>
          </Grid>

          <Grid item xs={12}>
            <Button
              component={RouterLink}
              to="/room_data_lists"
              variant="contained"
            >
              BACK
            </Button>
            <Button
              style={{ float: "right" }}
              variant="contained"
              onClick={submit}
              color="primary"
            >
              SUBMIT
            </Button>
          </Grid>
        </Grid>
      </Paper>
    </Container>
  );
}
export default RoomDataListCreate;
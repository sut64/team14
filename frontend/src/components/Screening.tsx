import { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import Typography from "@material-ui/core/Typography";
import Button from "@material-ui/core/Button";
import Container from "@material-ui/core/Container";
import Paper from "@material-ui/core/Paper";
import Box from "@material-ui/core/Box";
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableContainer from "@material-ui/core/TableContainer";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import { format } from 'date-fns'
import { ScreeningInterface } from "../models/IScreening";

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    container: {
      marginTop: theme.spacing(2),
    },
    table: {
      minWidth: 650,
    },
    tableSpace: {
      marginTop: 20,
    },
  })
);

function Screening() {
  const classes = useStyles();
  const [screening, setScreening] = useState<ScreeningInterface[]>([]);
  
  const getScreening = async () => {
    const apiUrl = "http://localhost:8080";
  const requestOptions = {
    method: "GET",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
    },
  };
    fetch(`${apiUrl}/screenings`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        console.log(res.data);
        if (res.data) {
          setScreening(res.data);
        } else {
          console.log("else");
        }
      });
  };
  
  useEffect(() => {
    getScreening();
  }, []);

  return (
    <div>
      <Container className={classes.container} maxWidth="md">
        <Box display="flex">
          <Box flexGrow={1}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              ผู้ป่วยทั้งหมด
            </Typography>
          </Box>
          <Box>
            <Button
              component={RouterLink}
              to="/screening/create"
              variant="contained"
              color="primary"
            >
              จัดการผู้ป่วย
            </Button>
          </Box>
        </Box>
        <TableContainer component={Paper} className={classes.tableSpace}>
          <Table className={classes.table} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="center" width="5%">
                  ลำดับ
                </TableCell>
                <TableCell align="center" width="15%">
                  ชื่อ
                </TableCell>
                <TableCell align="center" width="15%">
                  อาการ
                </TableCell>
                <TableCell align="center" width="25%">
                  จำนวนวันที่ป่วย
                </TableCell>
                <TableCell align="center" width="10%">
                  ห้องพัก
                </TableCell>
                <TableCell align="center" width="30%">
                  วันที่และเวลา
                </TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {screening.map((item: ScreeningInterface) => (
                <TableRow key={item.ID}>
                  <TableCell align="center">{item.ID}</TableCell>
                  <TableCell align="center">{item.Patient.Name}</TableCell>
                  <TableCell align="center">{item.Symptom.State}</TableCell>
                  <TableCell align="center">{item.Symptom.Period}</TableCell>
                  <TableCell align="center">{item.Room.RoomNumber}</TableCell>
                  <TableCell align="center">{format((new Date(item.Time)), 'dd MMMM yyyy hh:mm a')}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Container>
    </div>
  );
}

export default Screening;

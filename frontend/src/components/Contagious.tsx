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
import { ContagiousInterface } from "../models/IContagious";


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

function Contagious() {
  const classes = useStyles();
  const [contagious, setContagious] = useState<ContagiousInterface[]>([]);
  
  const apiUrl = "http://localhost:8080";
  const requestOptions = {
    method: "GET",
    headers: { 
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json" 
    },
  };

  
  const getContagious = async () => {
    fetch(`${apiUrl}/contagious`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        console.log(res.data);
        if (res.data) {
          setContagious(res.data);
        } else {
          console.log("else");
        }
      });
  };

  

  useEffect(() => {
    getContagious();
  }, []);  
  
  return (
    <div>
      <Container className={classes.container} maxWidth="lg">
        <Box display="flex">
          <Box flexGrow={1}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              ข้อมูลโรคติดต่อ
            </Typography>
          </Box>
          <Box >
            <Button
              component={RouterLink}
              to="/contagious/create"
              variant="contained"
              color="primary"
            >
              เพิ่มโรคติดต่อ
            </Button>
          </Box>
        </Box>
        <TableContainer component={Paper} className={classes.tableSpace}>
          <Table className={classes.table} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="center" width="2%">
                  ลำดับ
                </TableCell>
                <TableCell align="center" width="10%">
                  ชื่อโรคติดต่อ
                </TableCell>
                <TableCell align="center" width="10%">
                  เชื้อโรค
                </TableCell>
                <TableCell align="center" width="10%">
                  ประเภทการติดต่อ
                </TableCell>
                <TableCell align="center" width="20%">
                  อาการ
                </TableCell>
                <TableCell align="center" width="8%">
                  ระยะฟักตัวโดยเฉลี่ย (วัน)
                </TableCell>
                <TableCell align="center" width="10%">
                  ประเภทกลุ่มเสี่ยง
                </TableCell>
                <TableCell align="center" width="8%">
                  วันที่บันทึก
                </TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {contagious.map((item: ContagiousInterface) => (
                <TableRow key={item.ID}>
                  <TableCell align="center">{item.ID}</TableCell>
                  <TableCell align="center">{item.Name}</TableCell>
                  <TableCell align="center">{item.Germ.Name}</TableCell>
                  <TableCell align="center">{item.CatchingType.Title}</TableCell>
                  <TableCell align="center">{item.Symptom}</TableCell>
                  <TableCell align="center">{item.Incubation}</TableCell>
                  <TableCell align="center">{item.RiskGroupType.Title}</TableCell>
                  <TableCell align="center">{format((new Date(item.Date)), 'dd MM yyyy')}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Container>
    </div>
  );
}

export default Contagious;

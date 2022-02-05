import React, { useEffect, useState } from "react";
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
import { MedicineandVaccineInterface } from "../models/IMedicineandVaccine";
 
 
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
 
function MedicineandVaccines() {
 const classes = useStyles();
 const [medicineandvaccine, setMedicineandVaccines] = useState<MedicineandVaccineInterface[]>([]);
 const apiUrl = "http://localhost:8080";
  const requestOptions = {
    method: "GET",
    headers: { 
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json" },
  };
 const getMedicineandVaccines = async () => {
   fetch(`${apiUrl}/medicineand_vaccines`, requestOptions)
     .then((response) => response.json())
     .then((res) => {
       console.log(res.data);
       if (res.data) {
         setMedicineandVaccines(res.data);
       } else {
         console.log("else");
       }
     });
 };
 
 useEffect(() => {
   getMedicineandVaccines();
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
             ข้อมูลยาและวัคซีน
           </Typography>
         </Box>
         <Box>
           <Button
             component={RouterLink}
             to="/medicineand_vaccine/create"
             variant="contained"
             color="primary"
           >
             สร้างข้อมูล
           </Button>
         </Box>
       </Box>
       <TableContainer component={Paper} className={classes.tableSpace}>
         <Table className={classes.table} aria-label="simple table">
           <TableHead>
             <TableRow>
               <TableCell align="center" >
                 ID
               </TableCell>
               <TableCell align="center" >
                 RegNo
               </TableCell>
               <TableCell align="center" >
                 Name
               </TableCell>
               <TableCell align="center" >
                 Category
               </TableCell>
               <TableCell align="center" >
                 DosageForm
               </TableCell>
               <TableCell align="center" >
                 MinAge
               </TableCell>
               <TableCell align="center" >
                 MaxAge
               </TableCell>
               <TableCell align="center" >
                 Age
               </TableCell>
               <TableCell align="center" >
                 Contagious
               </TableCell>
               <TableCell align="center" >
                 Date
               </TableCell>
             </TableRow>
           </TableHead>
           <TableBody>
             {medicineandvaccine.map((item: MedicineandVaccineInterface) => (
               <TableRow key={item.ID}>
                 <TableCell align="center">{item.ID}</TableCell>
                 <TableCell align="center">{item.RegNo}</TableCell>
                 <TableCell align="center">{item.Name}</TableCell>
                 <TableCell align="center">{item.Category.Category}</TableCell>
                 <TableCell align="center">{item.DosageForm.DosageForm}</TableCell>
                 <TableCell align="center">{item.MinAge}</TableCell>
                 <TableCell align="center">{item.MaxAge}</TableCell>
                 <TableCell align="center">{item.Age.Age}</TableCell>
                 <TableCell align="center">{item.Contagious.Contagious}</TableCell>
                 <TableCell align="center">{format((new Date(item.Date)), 'dd MMMM yyyy hh:mm a')}</TableCell>
               </TableRow>
             ))}
           </TableBody>
         </Table>
       </TableContainer>
     </Container>
   </div>
 );
}
 
export default MedicineandVaccines;
 


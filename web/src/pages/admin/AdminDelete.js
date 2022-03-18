import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogContentText from '@mui/material/DialogContentText';
import DialogTitle from '@mui/material/DialogTitle';
import Button from '@mui/material/Button';
import {useSnackbar} from 'notistack';


import {adminDelete} from "@/api/api"


export function DeleteDialog(props) {
    const {enqueueSnackbar} = useSnackbar();
    const {open, handleClose, deleteId, deleteName} = props


    const handleDeleteAdmin = (id) => {
        (async () => {
            await adminDelete({id: id}).then((res) => {
                if (res.data.code === 0) {
                    enqueueSnackbar('Delete success!', {variant: 'success'});
                }
            })
            handleClose()
        })();
    }

    return (
        <Dialog
            fullWidth
            maxWidth="xs"
            open={open}
            onClose={handleClose}
            aria-labelledby="alert-dialog-title"
            aria-describedby="alert-dialog-description"
        >
            <DialogTitle id="alert-dialog-title">
                {"Delete Admin"}
            </DialogTitle>
            <DialogContent>
                <DialogContentText id="alert-dialog-description">
                    {"Are you sure you want to delete"}
                </DialogContentText>
                <DialogContentText id="alert-dialog-description">
                    <strong>{deleteName} ?</strong>
                </DialogContentText>
            </DialogContent>
            <DialogActions>
                <Button onClick={handleClose}>Cancel</Button>
                <Button onClick={() => {
                    handleDeleteAdmin(deleteId)

                }} autoFocus>
                    Delete Admin
                </Button>
            </DialogActions>
        </Dialog>
    )
}

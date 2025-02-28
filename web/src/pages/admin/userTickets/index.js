import React, { useEffect, useState } from 'react';
import { useParams, useNavigate, useLocation } from 'react-router-dom';
import { ReactComponent as IconMenu } from '../../../images/qr_code.svg';
import Ticket from '../../../components/ticket/ticket';
import { apiRequest } from '../../../services/api';
import alertMessage from '../../../components/alertMessage';
import Loading from '../../../components/loading';
import Header from '../../../components/header';
import { formatCpf } from '../../../utils';

export default function AdminUserTickets() {
  const { userId } = useParams();
  const [user, setUser] = useState([]);
  const [loading, setLoading] = useState(true);
  const location = useLocation();
  const navigate = useNavigate();

  useEffect(() => {
    reloadTickets();
  }, []);

  const reloadTickets = () => {
    setLoading(true);
    apiRequest('login', `login/v0/user/${userId}`, 'get')
      .then(async (response) => {
        const responseBody = await response.json();
        if (response.ok) {
          setUser(responseBody.data);
        } else {
          alertMessage('error', responseBody.error);
        }
      })
      .catch(() => {
        alertMessage('error', null);
      })
      .finally(() => {
        setLoading(false);
      });
  };

  if (loading) return <Loading />;

  return (
    <div>
      <Header title="Ingressos" />

      <div className="d-flex gap-3 justify-content-start mb-4">
        <div>
          <p className="fs-4 fw-bold text-white mb-1">{user.name}</p>
          <p className="text-white">CPF {formatCpf(user.cpf)}</p>
        </div>
        <div className="pt-2">
          <button
            className="btn btn-primary py-2 px-4"
            onClick={() =>
              navigate(`/admin/users/${userId}/ticket/create`, {
                state: { from: location }
              })
            }>
            <IconMenu className="me-2" />
            Criar ingresso
          </button>
        </div>
      </div>

      <div className="row gap-4 m-0">
        {!user.ticketlist || user.ticketlist.length === 0 ? (
          <div className="row justify-content-center align-items-center">
            <div className="col-12 fs-2 mb-4 text-center">Você ainda não possui ingressos</div>
            <div className="col-12" style={{ maxWidth: '350px' }}>
              <img
                src="/images/empty-tickets.svg"
                className="img-fluid"
                alt="Empty tickets image"
              />
            </div>
          </div>
        ) : (
          user.ticketlist.map(function (ticket) {
            return (
              <Ticket key={ticket.id} user={user} ticket={ticket} reloadTickets={reloadTickets} />
            );
          })
        )}
      </div>
    </div>
  );
}

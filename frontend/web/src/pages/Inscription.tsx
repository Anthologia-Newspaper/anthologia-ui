import * as React from 'react';
import { useEffect, useState } from 'react';
import { Button, Link } from '@chakra-ui/react';
import { Link as RouteLink, useNavigate } from 'react-router-dom';

import FormInput from 'components/FormInput';
import PwdInput from 'components/PwdInput';
import services from 'services';

const Inscription = (): JSX.Element => {
	const emailRegex = new RegExp(/^\w+([.-]?\w+)*@\w+([.-]?\w+)*(\.\w{2,3})+$/);

	const navigate = useNavigate();
	const [email, setEmail] = useState('');
	const [username, setUsername] = useState('');
	const [pwd, setPwd] = useState('');
	const [confirmedPwd, setConfirmedPwd] = useState('');
	const [validation, setValidation] = useState({
		valid: false,
		details: {
			email: { value: false, message: 'Email non renseigné.' },
			username: { value: false, message: "Nom d'utilisateur non renseigné." },
			pwd: { value: false, message: 'Mot de passe non renseigné.' },
			confirmedPwd: { value: false, message: 'Confirmation du mot de passe non renseignée.' },
		},
	});

	const inscription = async () => {
		try {
			const res = await services.authService.register({ email: 'a@a.com', password: 'aaa', username: 'aaa' });
			console.log(res);
			navigate('/favoris');
		} catch (error) {
			console.log(error);
		}
	};

	useEffect(() => {
		const emailValidation = emailRegex.test(email);
		const usernameValidation = username.length >= 3;
		const pwdValidation = pwd.length >= 5;
		const confirmedPwdValidation = pwd === confirmedPwd;
		const globalValidation = emailValidation && usernameValidation && pwdValidation && confirmedPwdValidation;

		setValidation({
			valid: globalValidation,
			details: {
				email: { value: emailValidation, message: emailValidation ? '' : 'Email non valide.' },
				username: { value: usernameValidation, message: usernameValidation ? '' : "Nom d'utilisateur trop court." },
				pwd: { value: pwdValidation, message: pwdValidation ? '' : 'Mot de passe trop court.' },
				confirmedPwd: {
					value: confirmedPwdValidation,
					message: confirmedPwdValidation ? '' : 'Mots de passe différents.',
				},
			},
		});
	}, [email, username, pwd, confirmedPwd]);

	return (
		<>
			<FormInput
				inputId="inscription-email-input"
				isError={email !== '' && !validation.details.email.value}
				errorMessage={validation.details.email.message}
				variant="primary-1"
				placeholder="e-mail"
				onChange={(e) => setEmail(e.target.value)}
			/>
			<FormInput
				inputId="inscription-username-input"
				isError={username !== '' && !validation.details.username.value}
				errorMessage={validation.details.username.message}
				variant="primary-1"
				placeholder="nom d'utilisateur"
				onChange={(e) => setUsername(e.target.value)}
			/>
			<PwdInput
				inputId="inscription-pwd-input"
				isError={pwd !== '' && !validation.details.pwd.value}
				errorMessage={validation.details.pwd.message}
				variant="primary-1"
				placeholder="mot de passe"
				onChange={(e) => setPwd(e.target.value)}
			/>
			<PwdInput
				inputId="inscription-confirmed-pwd-input"
				isError={confirmedPwd !== '' && !validation.details.confirmedPwd.value}
				errorMessage={validation.details.confirmedPwd.message}
				variant="primary-1"
				placeholder="confirmation du mot de passe"
				onChange={(e) => setConfirmedPwd(e.target.value)}
			/>
			<Button id="inscription-inscription-btn" variant="primary-1" isDisabled={!validation.valid} onClick={inscription}>
				Inscription
			</Button>
			<Link as={RouteLink} to="/connexion" w="100%">
				<Button id="inscription-connexion-btn" variant="secondary-4">
					Connexion
				</Button>
			</Link>
		</>
	);
};

export default Inscription;
